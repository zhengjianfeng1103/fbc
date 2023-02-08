package feesplit

import (
	"github.com/zhengjianfeng1103/fbc/x/feesplit/keeper"
	"github.com/zhengjianfeng1103/fbc/x/feesplit/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	NewKeeper           = keeper.NewKeeper
	SetParamsNeedUpdate = types.SetParamsNeedUpdate
)

type (
	Keeper = keeper.Keeper
)
