package getstatus

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/getstatus/getstatus_erc721"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/get-status")
	{
		getstatus_erc721.ApplyRoutes(g)
	}
}
