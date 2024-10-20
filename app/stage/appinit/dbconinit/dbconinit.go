// Package dbconinit provides method to Init database
package dbconinit

import (
	"fmt"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/Weareflexable/Superiad/config/envconfig"
	"github.com/Weareflexable/Superiad/pkg/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	var (
		host     = envconfig.EnvVars.DB_HOST
		username = envconfig.EnvVars.DB_USERNAME
		password = envconfig.EnvVars.DB_PASSWORD
		dbname   = envconfig.EnvVars.DB_NAME
		port     = envconfig.EnvVars.DB_PORT
		ssh      = envconfig.EnvVars.DB_SSL_MODE
	)

	if len(envconfig.EnvVars.DB_SSL_MODE) == 0 || envconfig.EnvVars.DB_SSL_MODE == "nil" {
		ssh = "disable"
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%d",
		host, username, password, dbname, ssh, port)

	var err error

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		logo.Fatal("failed to connect database", err)
	}

	// Store database in global store
	store.DB = db

	// Get underlying sql database to ping it
	sqlDb, err := db.DB()
	if err != nil {
		logo.Fatal("failed to ping database", err)
	}

	// If ping fails then log error and exit
	if err = sqlDb.Ping(); err != nil {
		logo.Fatal("failed to ping database", err)
	}
}
