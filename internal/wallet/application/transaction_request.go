package application

import (
	common_domain "github.com/paguerre3/blockchain/internal/common/domain"
)

type TransactionRequest struct {
	IdempotentKey              *string                  `json:"_idempotentKey"`
	SenderPrivateKey           *string                  `json:"senderPrivateKey"`
	SenderPublicKey            *string                  `json:"senderPublicKey"`
	SenderBlockChainAddress    *string                  `json:"senderBlockChainAddress"`
	RecipientBlockChainAddress *string                  `json:"recipientBlockChainAddress"`
	Amount                     *float64                 `json:"amount"`
	Signature                  *common_domain.Signature `json:"signature"`
}

func (t *TransactionRequest) ValidateIdempotency() error {
	// TODO
	return nil
}
