package projections

import (
	"context"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/domain"
	"wallet-api/domain/events"
)

const balanceProjectorType projector.Type = "balance"

type Balance struct {
	WalletID uuid.UUID `bson:"_id"`
	Version  int       `bson:"version"`
	Amount   int64     `bson:"amount"`
}

type BalanceProjector struct{}

func (b *Balance) EntityID() uuid.UUID {
	return b.WalletID
}

func (b *Balance) AggregateVersion() int {
	return b.Version
}

func NewBalanceProjector() *BalanceProjector {
	return &BalanceProjector{}
}

func (p *BalanceProjector) ProjectorType() projector.Type {
	return balanceProjectorType
}

func (p *BalanceProjector) Project(_ context.Context, event eventhorizon.Event, entity eventhorizon.Entity) (eventhorizon.Entity, error) {
	balance, ok := entity.(*Balance)
	if !ok {
		zap.S().Errorf("projection entity could not be handled: %v", entity)
		return nil, domain.ErrUnprocessable{Message: "Projection entity could not be handled"}
	}

	switch event.EventType() {
	case domain.WalletCreatedEvent:
		if err := applyCreatEvent(event, balance); err != nil {
			return nil, err
		}
	case domain.WalletCreditedEvent:
		if err := applyCreditEvent(event, balance); err != nil {
			return nil, err
		}
	case domain.WalletDebitedEvent:
		if err := applyDebitEvent(event, balance); err != nil {
			return nil, err
		}
	}
	balance.Version++
	return balance, nil
}

func applyCreatEvent(event eventhorizon.Event, balance *Balance) error {
	content, ok := event.Data().(*events.WalletCreatedContent)
	if !ok {
		zap.S().Errorf("wallet creation event content could not be handled by projector: %v", content)
		return domain.ErrUnprocessable{Message: "Create wallet event content could not be handled"}
	}
	balance.WalletID = event.AggregateID()
	return nil
}

func applyCreditEvent(event eventhorizon.Event, balance *Balance) error {
	content, ok := event.Data().(*events.WalletCreditedContent)
	if !ok {
		zap.S().Errorf("wallet credit event content could not be handled by projector: %v", content)
		return domain.ErrUnprocessable{Message: "Wallet credit event content could not be handled"}
	}
	balance.Amount += content.Amount
	return nil
}

func applyDebitEvent(event eventhorizon.Event, balance *Balance) error {
	content, ok := event.Data().(*events.WalletDebitedContent)
	if !ok {
		zap.S().Errorf("wallet debit event content could not be handled by projector: %v", content)
		return domain.ErrUnprocessable{Message: "Wallet debit event content could not be handled"}
	}
	balance.Amount -= content.Amount
	return nil
}
