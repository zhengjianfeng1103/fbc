package continuousauction

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"

	"github.com/zhengjianfeng1103/fbc/x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
