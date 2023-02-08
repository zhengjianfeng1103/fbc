//go:build ignore
// +build ignore

package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	swaptypes "github.com/zhengjianfeng1103/fbc/x/ammswap/types"
)

func TestInvariants(t *testing.T) {
	ctx, keeper := GetKeeper(t)
	keeper.swapKeeper.SetParams(ctx, swaptypes.DefaultParams())
	initPoolsAndLockInfos(t, ctx, keeper)

	_, broken := yieldFarmingAccountInvariant(keeper.Keeper)(ctx)
	require.False(t, broken)
	_, broken = moduleAccountInvariant(keeper.Keeper)(ctx)
	require.False(t, broken)
	_, broken = mintFarmingAccountInvariant(keeper.Keeper)(ctx)
	require.False(t, broken)
}
