package setstatus_erc721

type SetStatusRequest struct {
	TokenId         int64  `json:"tokenId" binding:"required"`
	ContractAddress string `json:"contractAddr" binding:"required"`
	Status          string `json:"status" binding:"required,min=1"`
}

type SetStatusPayload struct {
	TrasactionHash string
}
