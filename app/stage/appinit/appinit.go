package appinit

import (
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/dbconinit"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/dbmigrate"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/envinit"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/logoinit"
	"github.com/TheLazarusNetwork/superiad/pkg/platform"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
	platform.Init()
}
