package polygon

import (
	"context"
	"fmt"
	"math/big"

	"github.com/TheLazarusNetwork/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func Transfer(mnemonic string, to common.Address, value big.Int) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get wallet: %w", err)
		return "", err
	}
	publicKey := privKey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		err = fmt.Errorf("failed to dial client: %w", err)
		return "", err
	}

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		err = fmt.Errorf("failed to suggestGasTipCap, error %s", err)
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(int64(chainId)),
	}
	bn, err := client.BlockNumber(context.Background())
	if err != nil {
		err = fmt.Errorf("failed to get latest block no: %w", err)
		return "", err
	}

	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(publicKey), big.NewInt(int64(bn)))
	if err != nil {
		err = fmt.Errorf("failed to get nonce: %w", err)
		return "", err
	}

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, err := client.BlockByNumber(context.Background(), bignumBn)
	if err != nil {
		err = fmt.Errorf("failed to get latest block details: %w", err)
		return "", err
	}
	baseFee := misc.CalcBaseFee(config, blk.Header())
	gasMul := big.NewInt(3)
	mulRes := big.NewInt(0).Mul(baseFee, gasMul)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	maxFeePerGas = big.NewInt(80000000000)
	maxPriorityFeePerGas = big.NewInt(70000000000)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainId)),
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       21000,
		To:        &to,
		Value:     &value,
	})
	types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	if err != nil {
		err = fmt.Errorf("failed to sign tx: %w", err)
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		err = fmt.Errorf("failed to send tx: %w", err)
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
