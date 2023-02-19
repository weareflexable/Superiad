package polygon

import (
	"fmt"

	"github.com/Weareflexable/Superiad/config/envconfig"
	"github.com/Weareflexable/Superiad/generated/flexablenft"
	"github.com/Weareflexable/Superiad/pkg/wallet"
	rawtrasaction "github.com/Weareflexable/Superiad/pkg/wallet/rawtransaction"
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
