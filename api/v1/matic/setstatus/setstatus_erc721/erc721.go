package setstatus_erc721

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/config/envconfig"
	"github.com/TheLazarusNetwork/superiad/generated/generc721"
	"github.com/TheLazarusNetwork/superiad/models/contracts"
	"github.com/TheLazarusNetwork/superiad/pkg/network/polygon"
	"github.com/TheLazarusNetwork/superiad/pkg/platform"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc721")
	{

		g.POST("", setStatusErc721)
	}
}

func setStatusErc721(c *gin.Context) {
	network := "matic"
	var req SetStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	contractDetails, err := contracts.GetContract(req.ContractAddress)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(404, "contract not found in database").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch contract details from database").SendD(c)
		logo.Errorf("failed to fetch contract details from database, error: %s", err)
		return
	}
	if contractDetails.ContractType != contracts.ERC721 {
		httpo.NewErrorResponse(http.StatusBadRequest, "contract is not ERC721").SendD(c)
		return
	}

	erc721ContractAddr := common.HexToAddress(req.ContractAddress)

	bigIntTokenId := big.NewInt(req.TokenId)
	var hash string
	hash, err = polygon.SetStatus(erc721ContractAddr, req.Status, bigIntTokenId)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to set status").SendD(c)
		logo.Errorf("failed to set status of token with id: %d network: %s, contractAddr: %s, error: %s",
			req.TokenId, network, req.ContractAddress, err)
		return
	}
	payload := SetStatusPayload{
		TrasactionHash: hash,
	}

	httpo.NewSuccessResponseP(200, "trasaction initiated", payload).SendD(c)
}

func PlatformRoutine() {
	client, err := ethclient.Dial(polygon.GetRpcUrl())
	if err != nil {
		err = fmt.Errorf("failed to dial rpc url: %w", err)
		logo.Fatal(err)
	}
	contractAddr := common.HexToAddress(envconfig.EnvVars.FLEXABLE_CONTRACT_ADDRESS)
	MATIC_SETSTATUS_ERC721_ENDPOINT := "matic/set-status/erc721"
	instance, err := generc721.NewErc721(contractAddr, client)
	if err != nil {
		logo.Fatalf("failed to load contract at address %s, error: %s", contractAddr, err)
	}
	statusUpdatedChannel := make(chan *generc721.Erc721StatusUpdated, 10)
	_, err = instance.WatchStatusUpdated(nil, statusUpdatedChannel)
	if err != nil {
		logo.Fatalf("failed to listen to an event %s, error: %s", "StatusUpdated", err)
	}
	for {
		event := <-statusUpdatedChannel
		err = platform.TransactionHash(event.Raw.TxHash.Hex(), event.TokenID.Int64(), MATIC_SETSTATUS_ERC721_ENDPOINT)
		if err != nil {
			logo.Errorf("failed to report to platform: event: %s, error: %s", "StatusUpdated", err)
		}
	}
}
