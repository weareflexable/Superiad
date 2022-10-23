package delegate_erc721

type DelegateErc721Request struct {
	UserId          string `json:"userId" binding:"required"`
	ContractAddress string `json:"contractAddr" binding:"required"`
	MetaDataHash    string `json:"metaDataHash" binding:"required,min=1"`
}

type DelegateErc721Payload struct {
	TrasactionHash string
}
