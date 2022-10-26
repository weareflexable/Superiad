package platform

type TransactionSuccessRequest struct {
	TransactionHash string `json:"transactionHash" binding:"required,len=66"`
	TokenId         int64  `json:"tokenId" binding:"required"`
	EndPoint        string `json:"endPoint" binding:"required"`
}
