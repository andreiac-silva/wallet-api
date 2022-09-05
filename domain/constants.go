package domain

import eh "github.com/looplab/eventhorizon"

// WalletAggregateType represent the wallet aggregate type.
const WalletAggregateType eh.AggregateType = "wallet"

// Possible Commands.
const (
	WalletCreationCommand eh.CommandType = "create"
	WalletCreditCommand   eh.CommandType = "credit"
	WalletDebitCommand    eh.CommandType = "debit"
)

// Possible Events.
const (
	WalletCreatedEvent  eh.EventType = "wallet_created"
	WalletCreditedEvent eh.EventType = "amount_credited"
	WalletDebitedEvent  eh.EventType = "amount_debited"
)
