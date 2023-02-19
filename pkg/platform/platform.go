package platform

import (
	"github.com/Weareflexable/Superiad/config/envconfig"
	"github.com/go-resty/resty/v2"
)

var platformClient *resty.Client

func Init() {
	platformClient = resty.New().
		SetBaseURL(envconfig.EnvVars.PLATFORM_BASE_URL).
		SetHeader("Authorization", envconfig.EnvVars.PLATFORM_TOKEN)
}
