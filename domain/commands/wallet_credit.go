package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"

	"wallet-api/domain"
)

type Credit struct {
	ID          uuid.UUID `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description" eh:"optional"`
}

func NewCreditCommand(id uuid.UUID, amount int64, description string) Credit {
	return Credit{
		ID:          id,
		Amount:      amount,
		Description: description,
	}
}

func (c Credit) AggregateID() uuid.UUID {
	return c.ID
}

func (c Credit) AggregateType() eh.AggregateType {
	return domain.WalletAggregateType
}

func (c Credit) CommandType() eh.CommandType {
	return domain.WalletCreditCommand
}
