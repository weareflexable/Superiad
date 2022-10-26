package apiroutines

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/delegate/delegate_erc721"
	"github.com/TheLazarusNetwork/superiad/api/v1/matic/setstatus/setstatus_erc721"
)

func Run() {
	go setstatus_erc721.PlatformRoutine()
	go delegate_erc721.PlatformRoutine()
}
