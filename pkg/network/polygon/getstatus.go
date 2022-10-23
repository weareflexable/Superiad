package polygon

import (
	"math/big"

	"github.com/TheLazarusNetwork/superiad/generated/generc721"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetStatus(contractAddr common.Address, tokenId *big.Int) (string, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "nil", err
	}
	ins, err := generc721.NewErc721(contractAddr, client)
	if err != nil {
		return "nil", err
	}
	status, err := ins.FlexableNFTNFTs(nil, tokenId)
	if err != nil {
		return "", err
	} else {
		return status, nil
	}
}
