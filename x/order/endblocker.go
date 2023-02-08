package order

import (
	"fmt"

	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"

	"github.com/zhengjianfeng1103/fbc/x/common/perf"
	"github.com/zhengjianfeng1103/fbc/x/order/keeper"
	"github.com/zhengjianfeng1103/fbc/x/order/match"
	"github.com/zhengjianfeng1103/fbc/x/order/types"
)

// EndBlocker called every block
// 1. execute matching engine
// 2. flush cache
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {

	seq := perf.GetPerf().OnEndBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnEndBlockExit(ctx, types.ModuleName, seq)

	match.GetEngine().Run(ctx, keeper)

	// flush cache at the end
	keeper.Cache2Disk(ctx)

	keeper.SetMetric()
	ret := keeper.GetOperationMetric()

	tailmsg := func(name string, num int64) string {
		var msg string
		if num != 0 {
			msg = fmt.Sprintf("%s<%d>,", name, num)
		}
		return msg
	}

	message := tailmsg("FullFillNum", ret.FullFillNum)
	message += tailmsg("OpenNum", ret.OpenNum)
	message += tailmsg("CancelNum", ret.CancelNum)
	message += tailmsg("ExpireNum", ret.ExpireNum)
	message += tailmsg("PartialFillNum", ret.PartialFillNum)
	perf.GetPerf().EnqueueMsg(message)
}
