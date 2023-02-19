package delegatenftcreation

type DelegateTicketCreationRequest struct {
	UserId          string `json:"userId" binding:"required"`
	ContractAddress string `json:"contractAddr" binding:"required"`
	MetadataURI     string `json:"metadataURI" binding:"required,min=1"`
}

type DelegateTicketCreationPayload struct {
	TransactionHash string
}
