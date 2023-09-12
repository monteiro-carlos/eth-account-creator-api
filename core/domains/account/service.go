package account

import (
	"crypto/ecdsa"
	"encoding/json"
	"eth-account-creator-api/core/domains/account/models"
	"eth-account-creator-api/core/domains/adapters/queue"
	"eth-account-creator-api/internal/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

type ServiceI interface {
	CreateAddress() (models.Account, error)
	FetchAddressFromPubKey(publicKeyStr string) (models.Address, error)
	SendTransaction(transaction models.Transaction) error
}

type Account struct {
	log         *log.Logger
	queueClient *queue.Client
}

func NewAccountService(logger *log.Logger, queueClient *queue.Client) (*Account, error) {
	return &Account{
		log:         logger,
		queueClient: queueClient,
	}, nil
}

func (a *Account) CreateAddress() (models.Account, error) {
	account := models.Account{}

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		a.log.Zap.Fatal("error", zap.Error(err))
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		a.log.Zap.Fatal("Error casting public key to ECDSA", zap.Error(nil))
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressStr := address.Hex()

	account.PrivateKey = privateKeyStr
	account.PublicKey = publicKeyStr
	account.Address = addressStr

	return account, err
}

func (a *Account) FetchAddressFromPubKey(publicKeyStr string) (models.Address, error) {
	address := models.Address{}
	publicKeyBytes, err := hexutil.Decode(publicKeyStr)
	if err != nil {
		a.log.Zap.Fatal("error", zap.Error(err))
		return address, err
	}
	publicKeyECDSA, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		a.log.Zap.Fatal("Error unmarshalling public key", zap.Error(err))
	}

	addressObj := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressStr := addressObj.Hex()

	address.Address = addressStr

	return address, nil
}

func (a *Account) SendTransaction(transaction models.Transaction) error {
	var transactionMap map[string]interface{}
	data, err := json.Marshal(transaction)
	if err != nil {
		a.log.Zap.Fatal("error", zap.Error(err))
		return err
	}
	json.Unmarshal(data, &transactionMap)

	if err := a.queueClient.SendMessage(transactionMap); err != nil {
		a.log.Zap.Fatal("Error", zap.Error(err))

	}

	return err
}
