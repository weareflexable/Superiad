package matic

import (
	"github.com/Weareflexable/Superiad/api/middleware/onlyunlockedmiddleware"
	"github.com/Weareflexable/Superiad/api/v1/matic/approve"
	approveall "github.com/Weareflexable/Superiad/api/v1/matic/approveAll"
	"github.com/Weareflexable/Superiad/api/v1/matic/checkbalance"
	"github.com/Weareflexable/Superiad/api/v1/matic/delegateticketcreation/delegatenftcreation"
	"github.com/Weareflexable/Superiad/api/v1/matic/fetchwallet"
	"github.com/Weareflexable/Superiad/api/v1/matic/getstatus"
	"github.com/Weareflexable/Superiad/api/v1/matic/isowner"
	"github.com/Weareflexable/Superiad/api/v1/matic/setstatus"
	signmessage "github.com/Weareflexable/Superiad/api/v1/matic/signMessage"
	"github.com/Weareflexable/Superiad/api/v1/matic/tokenuri"
	"github.com/Weareflexable/Superiad/api/v1/matic/transfer"
	verifysignature "github.com/Weareflexable/Superiad/api/v1/matic/verifySignature"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
	{
		delegatenftcreation.ApplyRoutes(v1)

		checkbalance.ApplyRoutes(v1)
		fetchwallet.ApplyRoutes(v1)
		isowner.ApplyRoutes(v1)
		verifysignature.ApplyRoutes(v1)
		tokenuri.ApplyRoutes(v1)
		getstatus.ApplyRoutes(v1)
		setstatus.ApplyRoutes(v1)

		v1.Use(onlyunlockedmiddleware.OnlyUnlocked())
		signmessage.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		approve.ApplyRoutes(v1)
		approveall.ApplyRoutes(v1)
	}
}
