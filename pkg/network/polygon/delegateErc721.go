package polygon

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Weareflexable/Superiad/config/envconfig"
	"github.com/Weareflexable/Superiad/generated/flexablenft"
	"github.com/Weareflexable/Superiad/pkg/network/base/contract"
	"github.com/Weareflexable/Superiad/pkg/wallet"
	rawtrasaction "github.com/Weareflexable/Superiad/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DelegateNFTCreation(mnemonic string, contractAddr common.Address, metadataURI string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get user wallet from mnemonic: %w", err)
		return "", err
	}

	operatorPrivKey, err := wallet.GetWallet(envconfig.EnvVars.OPERATOR_MNEMONIC, GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get operator wallet from mnemonic: %w", err)
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		err = fmt.Errorf("failed to dial rpc url: %w", err)
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	tx, err := rawtrasaction.SendRawTrasac(operatorPrivKey, *client, int64(chainId), 310000, contractAddr, flexablenft.FlexablenftMetaData.ABI, "delegateTicketCreation", walletAddr, metadataURI)
	if err != nil {
		err = fmt.Errorf("failed to send raw transaction: %w", err)
		return "", err
	}
	return tx.Hash().Hex(), nil

}

func DelegateNFTCreationForBase(mnemonic string, metadataURI string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get user wallet from mnemonic: %w", err)
		return "", err
	}

	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	// bindTransactionDetails := BaseChainId()

	var contract contract.ContractTransactor

	fmt.Println("walletAddr, metadataURI : ", walletAddr, metadataURI)

	var test *bind.TransactOpts

	txn, err := contract.DelegateTicketCreation(test, walletAddr, metadataURI)

	fmt.Println("err : ", err.Error())

	return txn.ChainId().String(), err

}
func BaseChainId() *bind.TransactOpts {
	// Connect to Ethereum client

	BASE_RPC_URL := "https://sepolia.base.org/"

	var auth *bind.TransactOpts
	client, err := ethclient.Dial(BASE_RPC_URL)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
		return auth
	}

	privateKey, err := crypto.HexToECDSA(envconfig.EnvVars.PRIVATE_KEY)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load private key"})
		return auth

	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast public key to ECDSA"})
		return auth

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nonce"})
		return auth

	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gas price"})
		return auth

	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return auth

	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return auth

	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // Adjust as needed
	auth.GasPrice = gasPrice
	return auth

}
