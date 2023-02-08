package keeper

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/supply/exported"

	"github.com/zhengjianfeng1103/fbc/x/distribution/types"
)

// GetDistributionAccount returns the distribution ModuleAccount
func (k Keeper) GetDistributionAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}
