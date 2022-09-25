package aggregates

import (
	"context"
	"time"

	eh "github.com/looplab/eventhorizon"
	ehEvents "github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/domain"
	"wallet-api/domain/commands"
	"wallet-api/domain/events"
	"wallet-api/domain/usecases"
)

type Wallet struct {
	*ehEvents.AggregateBase

	balanceUc usecases.BalanceUseCase

	documentNumber string
	// Other aggregate fields...
}

func NewWallet(id uuid.UUID, balanceUc usecases.BalanceUseCase) *Wallet {
	return &Wallet{
		AggregateBase: ehEvents.NewAggregateBase(domain.WalletAggregateType, id),
		balanceUc:     balanceUc,
	}
}

func (w *Wallet) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case commands.CreateWallet:
		w.AppendEvent(
			domain.WalletCreatedEvent,
			events.NewWalletCreatedContent(cmd.DocumentNumber),
			time.Now().UTC(),
		)
		return nil
	case commands.Credit:
		w.AppendEvent(
			domain.WalletCreditedEvent,
			events.NewWalletCreditedContent(cmd.Amount, cmd.Description, cmd.ID.String()),
			time.Now().UTC(),
		)
		return nil
	case commands.Debit:
		// Validating debit attempts synchronously.
		if err := w.validateDebitAttempt(ctx, cmd.Amount); err != nil {
			return err
		}
		w.AppendEvent(
			domain.WalletDebitedEvent,
			events.NewWalletDebitedContent(cmd.Amount, cmd.Description, cmd.ID.String()),
			time.Now().UTC(),
		)
		return nil
	default:
		zap.S().Errorf("invalid command: %v", cmd)
		return domain.ErrUnprocessable{Message: "Command could not be processed"}
	}
}

func (w *Wallet) ApplyEvent(_ context.Context, e eh.Event) error {
	switch e.EventType() {
	case domain.WalletCreatedEvent:
		if data, ok := e.Data().(*events.WalletCreatedContent); ok {
			w.documentNumber = data.DocumentNumber
			return nil
		}
		zap.S().Errorf("invalid event data: %v", e.Data())
		return domain.ErrUnprocessable{Message: "Event data could not be processed"}
	default:
		return nil
	}
}

func (w *Wallet) validateDebitAttempt(ctx context.Context, amount int64) error {
	balance, err := w.balanceUc.CheckAvailableAmount(ctx, w.EntityID(), amount)
	if err != nil {
		zap.S().Error("debit attempt refused")
		return err
	}
	if balance.Version != w.AggregateVersion() {
		zap.S().Errorf(
			"incompatible projection version. projection version is %d, but the aggregate version is %d",
			balance.Version,
			w.AggregateVersion(),
		)
		return domain.ErrIncompatibleProjectionVersion{
			Message: "Operation refused because wallet balance is not up to date",
		}
	}
	return nil
}
