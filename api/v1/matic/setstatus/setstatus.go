package setstatus

import (
	"github.com/Weareflexable/Superiad/api/v1/matic/setstatus/setstatus_erc721"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/set-status")
	{
		setstatus_erc721.ApplyRoutes(g)
	}
}
