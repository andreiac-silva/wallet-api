package commands

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/uuid"

	"wallet-api/domain"
)

type CreateWallet struct {
	ID             uuid.UUID `json:"id"`
	DocumentNumber string    `json:"document_number"`
}

func NewCreateWalletCommand(id uuid.UUID, documentNumber string) *CreateWallet {
	return &CreateWallet{
		ID:             id,
		DocumentNumber: documentNumber,
	}
}

func (c CreateWallet) AggregateID() uuid.UUID {
	return c.ID
}

func (c CreateWallet) AggregateType() eh.AggregateType {
	return domain.WalletAggregateType
}

func (c CreateWallet) CommandType() eh.CommandType {
	return domain.WalletCreationCommand
}
