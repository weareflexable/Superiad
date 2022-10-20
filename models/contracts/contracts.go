package contracts

import "github.com/TheLazarusNetwork/superiad/pkg/store"

// Enum doesn't work properly for now
type Contract struct {
	Address      string       `gorm:"primary_key;not null"`
	ContractType ContractType `gorm:"contract_type" sql:"max_type"`
}

func GetContract(contractAddr string) (contract Contract, err error) {
	db := store.DB
	err = db.First(&contract, &Contract{
		Address: contractAddr,
	}).Error
	return
}
