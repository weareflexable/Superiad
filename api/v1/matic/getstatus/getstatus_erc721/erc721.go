package getstatus_erc721

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/models/contracts"
	"github.com/TheLazarusNetwork/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc721")
	{

		g.POST("", getStatusErc721)
	}
}

func getStatusErc721(c *gin.Context) {
	network := "matic"
	var req GetStatusRequest
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

	status, err := polygon.GetStatus(common.HexToAddress(req.ContractAddress), big.NewInt(req.TokenId))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get status of token").SendD(c)
		logo.Errorf("failed to get status of token with id: %d, network: %v, contractAddr: %v, error: %s",
			req.TokenId, network, req.ContractAddress, err)
		return
	}

	httpo.NewSuccessResponseP(200, "status fetched", GetStatusPayload{
		Status: status,
	}).SendD(c)

}
