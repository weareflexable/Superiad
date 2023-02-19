package platform

import (
	"fmt"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
)

func TransactionHash(transactionHash string, tokenId int64, endPoint string) error {
	var apiResSuccess httpo.ApiSuccessResponse[any]
	// var apiResFailure httpo.ApiErrorResponse[any]

	req := TransactionSuccessRequest{
		TransactionHash: transactionHash,
		TokenId:         tokenId,
		EndPoint:        endPoint,
	}
	res, err := platformClient.R().
		SetResult(&apiResSuccess).
		SetHeader("Content-Type", "application/json").
		SetBody(&req).Post(TRANSACTION_SUCCESS_ENDPOINT)
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return fmt.Errorf("%w: status code is %d, body is %s", ErrStatusCodeNotOk, res.StatusCode(), res.Body())
	}
	return nil
}
