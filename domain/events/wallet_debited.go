package events

type WalletDebitedContent struct {
	Amount      int64  `bson:"amount"`
	Description string `bson:"description"`
	WalletID    string `bson:"wallet_id"`
}

func NewWalletDebitedContent(amount int64, description, walletID string) *WalletDebitedContent {
	return &WalletDebitedContent{
		Amount:      amount,
		Description: description,
		WalletID:    walletID,
	}
}
