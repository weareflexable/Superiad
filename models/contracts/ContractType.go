package contracts

import (
	"database/sql/driver"
)

type ContractType string

const (
	ERC721 ContractType = "ERC721"
	ERC20  ContractType = "ERC20"
)

func (fit *ContractType) Scan(value interface{}) error {
	*fit = ContractType([]byte(value.(string)))
	return nil
}

func (fit ContractType) Value() (driver.Value, error) {
	return string(fit), nil
}
