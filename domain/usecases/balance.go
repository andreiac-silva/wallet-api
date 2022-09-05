package usecases

import (
	"context"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"
	"go.uber.org/zap"

	"wallet-api/domain/errors"
	"wallet-api/domain/projections"
)

type balanceUseCase struct {
	repo eh.ReadRepo
}

type BalanceUseCase interface {
	CheckAvailableAmount(ctx context.Context, id uuid.UUID, amount int64) (*projections.Balance, error)
}

func NewBalanceUseCase(balanceRepo eh.ReadRepo) BalanceUseCase {
	return &balanceUseCase{repo: balanceRepo}
}

func (b balanceUseCase) CheckAvailableAmount(ctx context.Context, id uuid.UUID, amount int64) (*projections.Balance, error) {
	balance, err := b.findOne(ctx, id)
	if err != nil {
		return nil, err
	}
	if balance.Amount < amount {
		zap.S().Errorf("insufficient amount on the wallet with id %s", id)
		return nil, internalErrors.ErrInsufficientAmount{}
	}
	return balance, nil
}

func (b balanceUseCase) findOne(ctx context.Context, id uuid.UUID) (*projections.Balance, error) {
	found, err := b.repo.Find(ctx, id)
	// TODO: Improve error handling.
	if found == nil || err != nil {
		zap.S().Errorf("balance with id %s not found", id.String())
		return nil, internalErrors.ErrNotFound{Message: "Balance not found"}
	}
	balance, ok := found.(*projections.Balance)
	if !ok {
		zap.S().Errorf("balance could not be handled by usecase: %v", found)
		return nil, &internalErrors.ErrUnprocessable{Message: "Unprocessable Balance"}
	}
	return balance, nil
}
