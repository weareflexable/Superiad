package app

import (
	"github.com/Weareflexable/Superiad/app/stage/appinit"
	"github.com/Weareflexable/Superiad/app/stage/apprun"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func InitRun() {
	appinit.Init()
	apprun.Run()
}
