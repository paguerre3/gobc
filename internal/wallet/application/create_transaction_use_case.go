package application

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/blockchain/configs"
	common_domain "github.com/paguerre3/blockchain/internal/common/domain"
	"github.com/paguerre3/blockchain/internal/wallet/domain"
)

var (
	txRequestLock = common_domain.NewLock()
	config        = configs.Instance()
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
	idempotencyKey := common_domain.ToSafeStr(transactionRequest.IdempotencyKey)
	lockTimeout := time.Duration(config.Lock().TimeOutInSeconds()) * time.Second
	if err := txRequestLock.Acquire(idempotencyKey, lockTimeout); err != nil {
		return fmt.Errorf("failed to acquire lock for key %s wiyh request %+v: %w",
			idempotencyKey, transactionRequest, err)
	}
	defer txRequestLock.Release(idempotencyKey)

	senderPublicKey := common_domain.PublicKeyFromString(
		common_domain.ToSafeStr(transactionRequest.SenderPublicKey))
	senderPrivateKey := common_domain.PrivateKeyFromString(
		common_domain.ToSafeStr(transactionRequest.SenderPrivateKey),
		senderPublicKey)
	senderBlockChainAddress := common_domain.ToSafeStr(transactionRequest.SenderBlockChainAddress)
	recipientBlockChainAddress := common_domain.ToSafeStr(transactionRequest.RecipientBlockChainAddress)

	amount := common_domain.ToSafeFloat64(transactionRequest.Amount)

	newTransaction, err := domain.NewTransaction(senderPrivateKey, senderBlockChainAddress, recipientBlockChainAddress, amount)
	log.Infof("New Transaction: %+v", newTransaction) // including timestamp

	// TODO domain error in echo middleware ??
	return err
}
