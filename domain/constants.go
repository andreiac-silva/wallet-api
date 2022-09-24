package domain

import eh "github.com/looplab/eventhorizon"

// WalletAggregateType represent the wallet aggregate type.
const WalletAggregateType eh.AggregateType = "wallet"

// Commands.
const (
	WalletCreationCommand eh.CommandType = "create_wallet"
	WalletCreditCommand   eh.CommandType = "credit_amount"
	WalletDebitCommand    eh.CommandType = "debit_amount"
)

// Events.
const (
	WalletCreatedEvent  eh.EventType = "wallet_created"
	WalletCreditedEvent eh.EventType = "amount_credited"
	WalletDebitedEvent  eh.EventType = "amount_debited"
)
