package tokenuri

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/tokenuri/tokenuri_erc721"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/token-uri")
	{
		tokenuri_erc721.ApplyRoutes(g)
	}
}
