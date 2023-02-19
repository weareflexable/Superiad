package delegate

import (
	"github.com/Weareflexable/Superiad/api/v1/matic/delegateticketcreation/delegatenftcreation"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegate")
	{

		delegatenftcreation.ApplyRoutes(g)
	}
}
