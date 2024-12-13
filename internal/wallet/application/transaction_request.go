package application

import (
	common_domain "github.com/paguerre3/blockchain/internal/common/domain"
)

type TransactionRequest struct {
	IdempotencyKey             *string                  `json:"_idempotencyKey"`
	SenderPrivateKey           *string                  `json:"senderPrivateKey"`
	SenderPublicKey            *string                  `json:"senderPublicKey"`
	SenderBlockChainAddress    *string                  `json:"senderBlockChainAddress"`
	RecipientBlockChainAddress *string                  `json:"recipientBlockChainAddress"`
	Amount                     *float64                 `json:"amount"`
	Signature                  *common_domain.Signature `json:"signature"`
}
