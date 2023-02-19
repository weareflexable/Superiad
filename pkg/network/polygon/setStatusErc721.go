package polygon

import (
	"fmt"
	"math/big"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/Weareflexable/Superiad/config/envconfig"
	"github.com/Weareflexable/Superiad/generated/generc721"
	"github.com/Weareflexable/Superiad/pkg/wallet"
	rawtrasaction "github.com/Weareflexable/Superiad/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetStatus(contractAddr common.Address, status string, tokenId *big.Int) (string, error) {
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
	tx, err := rawtrasaction.SendRawTrasac(operatorPrivKey, *client, int64(chainId), 310000, contractAddr, generc721.Erc721MetaData.ABI, "setStatus", tokenId, status)
	if err != nil {
		err = fmt.Errorf("failed to send raw transaction: %w", err)
		return "", err
	}

	instance, err := generc721.NewErc721(contractAddr, client)
	if err != nil {
		logo.Errorf("failed to load contract at address %s, error: %s", contractAddr, err)
		return "", err
	}
	statusUpdatedChannel := make(chan *generc721.Erc721StatusUpdated, 10)
	subs, err := instance.WatchStatusUpdated(nil, statusUpdatedChannel)
	if err != nil {
		logo.Fatalf("failed to listen to an event %s, error: %s", "StatusUpdated", err)
		return "", err
	}

	subs.Unsubscribe()

	return tx.Hash().Hex(), nil

}
