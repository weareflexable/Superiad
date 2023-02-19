// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/Weareflexable/Superiad/models/contracts"
	"github.com/Weareflexable/Superiad/models/transaction"
	"github.com/Weareflexable/Superiad/models/user"
	"github.com/Weareflexable/Superiad/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &transaction.Transaction{}, &contracts.Contract{})
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
