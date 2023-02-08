package transfer

import (
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/apps/transfer/keeper"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/apps/transfer/types"
)

var (
	NewKeeper  = keeper.NewKeeper
	ModuleCdc  = types.ModuleCdc
	SetMarshal = types.SetMarshal
)
