package events

type WalletCreatedContent struct {
	DocumentNumber string `bson:"document_number"`
}

func NewWalletCreatedContent(documentNumber string) *WalletCreatedContent {
	return &WalletCreatedContent{
		DocumentNumber: documentNumber,
	}
}
