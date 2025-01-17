package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	APP_PORT                  int      `env:"APP_PORT,notEmpty"`
	GIN_MODE                  string   `env:"GIN_MODE,notEmpty"`
	DB_HOST                   string   `env:"DB_HOST,notEmpty"`
	DB_USERNAME               string   `env:"DB_USERNAME,notEmpty"`
	DB_PASSWORD               string   `env:"DB_PASSWORD,notEmpty"`
	DB_NAME                   string   `env:"DB_NAME,notEmpty"`
	DB_PORT                   int      `env:"DB_PORT,notEmpty"`
	ALLOWED_ORIGIN            []string `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	NETWORK_RPC_URL_ETHEREUM  string   `env:"NETWORK_RPC_URL_ETHEREUM,notEmpty"`
	NETWORK_RPC_URL_POLYGON   string   `env:"NETWORK_RPC_URL_POLYGON,notEmpty"`
	TOKEN                     string   `env:"TOKEN,notEmpty"`
	APP_ENVIRONMENT           string   `env:"APP_ENVIRONMENT,notEmpty"`
	OPERATOR_MNEMONIC         string   `env:"OPERATOR_MNEMONIC,notEmpty"`
	PLATFORM_BASE_URL         string   `env:"PLATFORM_BASE_URL,notEmpty"`
	PLATFORM_TOKEN            string   `env:"PLATFORM_TOKEN,notEmpty"`
	FLEXABLE_CONTRACT_ADDRESS string   `env:"FLEXABLE_CONTRACT_ADDRESS,notEmpty"`
	PRIVATE_KEY               string   `env:"PRIVATE_KEY,notEmpty"`
	DB_SSL_MODE               string   `env:"DB_SSL_MODE"`
}

var EnvVars config = config{}

func InitEnvVars() {
	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
	gin.SetMode(EnvVars.GIN_MODE)
}
