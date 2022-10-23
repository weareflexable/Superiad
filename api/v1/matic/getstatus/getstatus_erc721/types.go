package getstatus_erc721

type GetStatusRequest struct {
	TokenId         int64  `json:"tokenId" binding:"required"`
	ContractAddress string `json:"contractAddr" binding:"required"`
}

type GetStatusPayload struct {
	Status string
}
