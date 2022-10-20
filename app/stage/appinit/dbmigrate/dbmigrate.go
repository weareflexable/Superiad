// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/models/contracts"
	"github.com/TheLazarusNetwork/superiad/models/transaction"
	"github.com/TheLazarusNetwork/superiad/models/user"
	"github.com/TheLazarusNetwork/superiad/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &transaction.Transaction{}, &contracts.Contract{})
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
