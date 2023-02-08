package order

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"

	"github.com/zhengjianfeng1103/fbc/x/common/perf"
	"github.com/zhengjianfeng1103/fbc/x/order/keeper"
	"github.com/zhengjianfeng1103/fbc/x/order/types"
	//"github.com/zhengjianfeng1103/fbc/x/common/version"
)

// BeginBlocker runs the logic of BeginBlocker with version 0.
// BeginBlocker resets keeper cache.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
