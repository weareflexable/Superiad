package delegatenftcreation

import (
	"errors"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/Weareflexable/Superiad/models/contracts"
	"github.com/Weareflexable/Superiad/models/user"
	"github.com/Weareflexable/Superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateTicketCreation")
	{

		g.POST("", delegateTicketCreation)
	}
}

func delegateTicketCreation(c *gin.Context) {
	network := "matic"
	var req DelegateTicketCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)

			return
		}

		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logo.Errorf("failed to fetch user mnemonic for userId: %v, error: %s",
			req.UserId, err)
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

	flexableNFTContractAddr := common.HexToAddress(req.ContractAddress)
	var hash string
	hash, err = polygon.DelegateNFTCreation(mnemonic, flexableNFTContractAddr, req.MetadataURI)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to delegate erc721 to wallet of userId: %v , network: %v, contractAddr: %v, error: %s",
			req.UserId, network, req.ContractAddress, err)
		return
	}
	sendSuccessResponse(c, hash, req.UserId)
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := DelegateTicketCreationPayload{
		TransactionHash: hash,
	}
	if err := user.AddTrasactionHash(userId, hash); err != nil {
		logo.Errorf("failed to add transaction hash: %v to user id: %v, error: %s", hash, userId, err)
	}
	httpo.NewSuccessResponseP(200, "trasaction initiated", payload).SendD(c)
}

// func PlatformRoutine() {
// 	client, err := ethclient.Dial(polygon.GetRpcUrl())
// 	if err != nil {
// 		err = fmt.Errorf("failed to dial rpc url: %w", err)
// 		logo.Fatal(err)
// 	}
// 	contractAddr := common.HexToAddress(envconfig.EnvVars.FLEXABLE_CONTRACT_ADDRESS)

// 	MATIC_DELEGATE_ERC721_ENDPOINT := "matic/delegate/erc721"
// 	instance, err := generc721.NewErc721(contractAddr, client)
// 	if err != nil {
// 		logo.Fatalf("failed to load contract at address %s, error: %s", contractAddr, err)
// 	}
// 	ticketCreatedChannel := make(chan *generc721.Erc721TicketCreated, 10)
// 	_, err = instance.WatchTicketCreated(nil, ticketCreatedChannel, nil)
// 	if err != nil {
// 		logo.Fatalf("failed to listen to an event %s, error: %s", "ArtifactCreated", err)
// 	}
// 	for {
// 		event := <-ticketCreatedChannel
// 		err = platform.TransactionHash(event.Raw.TxHash.Hex(), event.TokenID.Int64(), MATIC_DELEGATE_ERC721_ENDPOINT)
// 		if err != nil {
// 			logo.Errorf("failed to report to platform: event: %s, error: %s", "ArtifactCreated", err)
// 		}
// 	}
// }
