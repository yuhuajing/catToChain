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

func gentx(gaslimit uint64, typechain int) *bind.TransactOpts {
	var privateKey *ecdsa.PrivateKey
	var err error
	switch typechain {
	case 0:
		//Prod
		//0xa39982a30d9DD7639F4340b3d90C6792cB38EE61
		privateKey, err = crypto.HexToECDSA("55cae8779390b54ce6b4328e47082fc5fb6bd1293b477c7e934853edd7884cf0")
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		//Logis
		//0x319E2178e70f28804946Ba4b178F8664f6AC0d0f
		privateKey, err = crypto.HexToECDSA("79c46a339c04777a5bd370aebaf124e2a14431effc270f60e9da44eebcf783a2")
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		//Sale
		//0x9d905f4c261c60c541355D9C2E3e0A941E3dCF67
		privateKey, err = crypto.HexToECDSA("dc8722425f3d307eccef3bbe55d27a25ea6c735b1fc2fb74ce223101d5a3c75c")
		if err != nil {
			log.Fatal(err)
		}
	case 3:
		//Repair
		//0x0DB76B5a2e462c9b2577E880F22762fE8AbF5819
		privateKey, err = crypto.HexToECDSA("b5daee2ba58d114be736fe4e20adfa637506f7790372645d322b856f3edabac6")
		if err != nil {
			log.Fatal(err)
		}
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
	auth := gentx(3000000, typechain)
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
