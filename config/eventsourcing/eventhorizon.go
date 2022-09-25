package eventsourcing

import (
	"context"
	"wallet-api/domain"

	"github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	mongoOutbox "github.com/looplab/eventhorizon/outbox/mongodb"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/domain/aggregates"
	"wallet-api/domain/events"
	"wallet-api/domain/projections"
	"wallet-api/domain/usecases"
)

func Setup(
	ctx context.Context,
	eventBus eventhorizon.EventBus,
	eventStore eventhorizon.EventStore,
	outbox *mongoOutbox.Outbox,
	commandBus *bus.CommandHandler,
	balanceRepo eventhorizon.ReadWriteRepo,
	balanceUc usecases.BalanceUseCase,
) {
	registerEvents()
	registerAggregates(balanceUc)

	// Add the Event Bus as the last handler of the outbox.
	if err := outbox.AddHandler(ctx, eventhorizon.MatchAll{}, eventBus); err != nil {
		zap.S().Fatal("could not add event bus to outbox:", err)
	}

	// Add a logger as an observer.
	if err := eventBus.AddHandler(ctx, eventhorizon.MatchAll{},
		eventhorizon.UseEventHandlerMiddleware(&Logger{})); err != nil {
		zap.S().Errorw("failure to add logger as an observer", "error", err)
	}

	// Create the aggregate store.
	aggregateStore, err := ehEvents.NewAggregateStore(eventStore)
	if err != nil {
		zap.S().Fatal("failure to create aggregate store", "error", err)
	}

	walletHandler, err := aggregate.NewCommandHandler(domain.WalletAggregateType, aggregateStore)
	if err != nil {
		zap.S().Fatal("failure to create command handler", "error", err)
	}

	if err := addCommandsToHandler(walletHandler, commandBus); err != nil {
		zap.S().Fatal("could not add command handler", "error", err)
	}

	// Configuring balance projection handler.
	balanceProjector := projector.NewEventHandler(projections.NewBalanceProjector(), balanceRepo)
	balanceProjector.SetEntityFactory(func() eventhorizon.Entity { return &projections.Balance{} })

	// Create and register a read model for a balance.
	err = eventBus.AddHandler(ctx, eventhorizon.MatchEvents{
		domain.WalletCreatedEvent,
		domain.WalletCreditedEvent,
		domain.WalletDebitedEvent,
	}, balanceProjector)

	if err != nil {
		zap.S().Fatal("could not create and register a read model for a balance", "error", err)
	}

	outbox.Start()

	zap.S().Debug("event horizon setup has done")
}

func registerEvents() {
	eventhorizon.RegisterEventData(domain.WalletCreatedEvent, func() eventhorizon.EventData {
		return &events.WalletCreatedContent{}
	})
	eventhorizon.RegisterEventData(domain.WalletCreditedEvent, func() eventhorizon.EventData {
		return &events.WalletCreditedContent{}
	})
	eventhorizon.RegisterEventData(domain.WalletDebitedEvent, func() eventhorizon.EventData {
		return new(events.WalletDebitedContent)
	})
}

func registerAggregates(uc usecases.BalanceUseCase) {
	eventhorizon.RegisterAggregate(func(id uuid.UUID) eventhorizon.Aggregate {
		return aggregates.NewWallet(id, uc)
	})
}

func addCommandsToHandler(walletHandler *aggregate.CommandHandler, commandBus *bus.CommandHandler) error {
	commandHandler := eventhorizon.UseCommandHandlerMiddleware(walletHandler, LoggingMiddleware)
	if err := commandBus.SetHandler(commandHandler, domain.WalletCreationCommand); err != nil {
		return err
	}
	if err := commandBus.SetHandler(commandHandler, domain.WalletCreditCommand); err != nil {
		return err
	}
	if err := commandBus.SetHandler(commandHandler, domain.WalletDebitCommand); err != nil {
		return err
	}
	return nil
}

func Close(
	eventBus eventhorizon.EventBus,
	eventStore eventhorizon.EventStore,
	outbox *mongoOutbox.Outbox,
	balanceRepo eventhorizon.ReadWriteRepo,
) {
	if err := balanceRepo.Close(); err != nil {
		zap.S().Errorw("error closing balance repo", "error", err)
	}
	if err := eventStore.Close(); err != nil {
		zap.S().Errorw("error closing event store", "error", err)
	}
	if err := outbox.Close(); err != nil {
		zap.S().Errorw("error closing outbox", "error", err)
	}
	if err := eventBus.Close(); err != nil {
		zap.S().Errorw("error closing event bus", "error", err)
	}
}
