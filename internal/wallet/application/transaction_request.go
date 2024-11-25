package application

type TransactionRequest struct {
	SenderPrivateKey           *string  `json:"senderPrivateKey"`
	SenderPublicKey            *string  `json:"senderPublicKey"`
	SenderBlockChainAddress    *string  `json:"senderBlockChainAddress"`
	RecipientBlockChainAddress *string  `json:"recipientBlockChainAddress"`
	Amount                     *float64 `json:"amount"`
}
