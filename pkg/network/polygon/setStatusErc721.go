package polygon

import (
	"fmt"
	"math/big"

	"github.com/TheLazarusNetwork/superiad/config/envconfig"
	"github.com/TheLazarusNetwork/superiad/generated/generc721"
	"github.com/TheLazarusNetwork/superiad/pkg/wallet"
	rawtrasaction "github.com/TheLazarusNetwork/superiad/pkg/wallet/rawtransaction"
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
	return tx.Hash().Hex(), nil

}
