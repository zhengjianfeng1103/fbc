package token

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/x/common/perf"
	"github.com/zhengjianfeng1103/fbc/x/token/types"
)

// BeginBlocker is called when dapp handles with abci::BeginBlock
func beginBlocker(ctx sdk.Context, keeper Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
