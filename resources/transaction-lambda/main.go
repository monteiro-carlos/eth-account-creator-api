package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EthereumTransaction(
	amountToTransfer float64, addressFromStr string, addressToStr string,
	privateKeyStr string, gasPrice uint64, gasLimit uint64,
) error {
	client, err := ethclient.Dial(os.Getenv("ETH_CLIENT_DIAL"))
	if err != nil {
		log.Fatal(err)
	}

	amount := new(big.Float)
	amount.SetFloat64(amountToTransfer)
	oneEthWei := new(big.Float)
	oneEthWei.SetInt(big.NewInt(int64(1000000000000000000)))
	amount.Mul(amount, oneEthWei)
	finalAmount := new(big.Int)
	f, _ := amount.Uint64()
	finalAmount.SetUint64(f)

	gasPriceFinal := new(big.Int)
	gasPriceFinal.SetUint64(gasPrice)

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	addressFrom := common.HexToAddress(addressFromStr)
	addressTo := common.HexToAddress(addressToStr)

	nonce, err := client.PendingNonceAt(context.Background(), addressFrom)
	if err != nil {
		log.Fatal(err)
	}

	transaction := types.NewTransaction(nonce, addressTo, finalAmount, gasLimit, gasPriceFinal, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal(err)
	}

	return err
}

func handler(sqsEvent events.SQSEvent) {
	fmt.Println(sqsEvent.Records[0].Body)
}

func main() {
	lambda.Start(handler)
}
