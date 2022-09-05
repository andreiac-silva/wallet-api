package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"

	"wallet-api/domain"
)

type Debit struct {
	ID          uuid.UUID `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description" eh:"optional"`
}

func NewDebitCommand(id uuid.UUID, amount int64, description string) Debit {
	return Debit{
		ID:          id,
		Amount:      amount,
		Description: description,
	}
}

func (d Debit) AggregateID() uuid.UUID {
	return d.ID
}

func (d Debit) AggregateType() eh.AggregateType {
	return domain.WalletAggregateType
}

func (d Debit) CommandType() eh.CommandType {
	return domain.WalletDebitCommand
}
