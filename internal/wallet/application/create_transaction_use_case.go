package application

import (
	common_domain "github.com/paguerre3/blockchain/internal/common/domain"
	"github.com/paguerre3/blockchain/internal/wallet/domain"
)

type CreateTransactionUseCase interface {
	Execute(transactionRequest TransactionRequest) error
}

type createTransactionUseCase struct {
	getWalletUseCase GetWalletUseCase
}

func NewCreateTransactionUseCase(getWalletUseCase GetWalletUseCase) CreateTransactionUseCase {
	return &createTransactionUseCase{
		getWalletUseCase: getWalletUseCase,
	}
}

func (c *createTransactionUseCase) Execute(transactionRequest TransactionRequest) error {
	// TODO
	transactionRequest.ValidateIdempotency()

	senderPublicKey := common_domain.PublicKeyFromString(
		common_domain.ToSafeStr(transactionRequest.SenderPublicKey))
	senderPrivateKey := common_domain.PrivateKeyFromString(
		common_domain.ToSafeStr(transactionRequest.SenderPrivateKey),
		senderPublicKey)
	senderBlockChainAddress := common_domain.ToSafeStr(transactionRequest.SenderBlockChainAddress)
	recipientBlockChainAddress := common_domain.ToSafeStr(transactionRequest.RecipientBlockChainAddress)

	amount := common_domain.ToSafeFloat64(transactionRequest.Amount)

	// TODO (add domain error)
	_, err := domain.NewTransaction(senderPrivateKey, senderBlockChainAddress, recipientBlockChainAddress, amount)
	return err
}
