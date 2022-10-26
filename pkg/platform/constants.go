package platform

import (
	"errors"
)

const (
	TRANSACTION_SUCCESS_ENDPOINT = "transaction-success"
)

var (
	ErrStatusCodeNotOk = errors.New("status code is not 200")
)
