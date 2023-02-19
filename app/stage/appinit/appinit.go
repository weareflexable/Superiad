package appinit

import (
	"github.com/Weareflexable/Superiad/app/stage/appinit/dbconinit"
	"github.com/Weareflexable/Superiad/app/stage/appinit/dbmigrate"
	"github.com/Weareflexable/Superiad/app/stage/appinit/envinit"
	"github.com/Weareflexable/Superiad/app/stage/appinit/logoinit"
	"github.com/Weareflexable/Superiad/pkg/platform"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
	platform.Init()
}
