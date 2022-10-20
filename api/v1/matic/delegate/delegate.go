package delegate

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/delegate/delegate_erc721"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegate")
	{

		delegate_erc721.ApplyRoutes(g)
	}
}
