package tokenuri_erc721

type TokenUriRequest struct {
	TokenId         int64  `json:"tokenId" binding:"required"`
	ContractAddress string `json:"contractAddr" binding:"required"`
}

type TokenUriPayload struct {
	TokenUri string `json:"tokenUri"`
}

