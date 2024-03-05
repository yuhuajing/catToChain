package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/trace"
	"math/big"
)

func gentx(gaslimit uint64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("57f40737d48920fd7e98c449f6aaa8498140ae91cc350d05e7c53bf02c6e20a1")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrUpdateData(number, description string, typechain int) string {
	address := common.HexToAddress(config.TraceSC)
	fmt.Println(1)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(3000000)
	switch typechain {
	case 0:
		tx, err := instance.AddProdData(auth, number, description)
		if err != nil {
			fmt.Println("error creating instance")
			log.Fatal(err)
		}
		return tx.Hash().Hex()
	case 1:
		tx, err := instance.AddLogisData(auth, number, description)
		if err != nil {
			fmt.Println("error creating instance")
			log.Fatal(err)
		}
		return tx.Hash().Hex()
	case 2:
		tx, err := instance.AddSaleData(auth, number, description)
		if err != nil {
			fmt.Println("error creating instance")
			log.Fatal(err)
		}
		return tx.Hash().Hex()
	case 3:
		tx, err := instance.AddRepairData(auth, number, description)
		if err != nil {
			fmt.Println("error creating instance")
			log.Fatal(err)
		}
		return tx.Hash().Hex()
	}
	return ""
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
