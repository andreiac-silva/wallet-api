package events

type WalletCreditedContent struct {
	Amount      int64  `bson:"amount"`
	Description string `bson:"description"`
	WalletID    string `bson:"wallet_id"`
}

func NewWalletCreditedContent(amount int64, description, walletID string) *WalletCreditedContent {
	return &WalletCreditedContent{
		Amount:      amount,
		Description: description,
		WalletID:    walletID,
	}
}
