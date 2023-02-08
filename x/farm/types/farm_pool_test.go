//go:build ignore
// +build ignore

package types

import (
	"testing"

	"github.com/zhengjianfeng1103/fbc/x/common"

	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

func TestFarmPools(t *testing.T) {
	tests := []struct {
		owner                   sdk.AccAddress
		name                    string
		lockedSymbol            string
		depositAmount           sdk.SysCoin
		totalValueLocked        sdk.SysCoin
		yieldedTokenInfos       YieldedTokenInfos
		totalAccumulatedRewards sdk.SysCoins
		isFinished              bool
	}{
		{
			owner:            sdk.AccAddress{0x1},
			name:             "pool",
			lockedSymbol:     "xxb",
			depositAmount:    sdk.NewDecCoin(common.NativeToken, sdk.ZeroInt()),
			totalValueLocked: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(100)),
			yieldedTokenInfos: YieldedTokenInfos{
				{
					RemainingAmount: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(100)),
				},
			},
			totalAccumulatedRewards: sdk.SysCoins{},
			isFinished:              false,
		},
		{
			owner:            sdk.AccAddress{0x1},
			name:             "pool",
			lockedSymbol:     "xxb",
			depositAmount:    sdk.NewDecCoin(common.NativeToken, sdk.ZeroInt()),
			totalValueLocked: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(0)),
			yieldedTokenInfos: YieldedTokenInfos{
				{
					RemainingAmount: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(100)),
				},
			},
			totalAccumulatedRewards: sdk.SysCoins{},
			isFinished:              false,
		},
		{
			owner:            sdk.AccAddress{0x1},
			name:             "pool",
			lockedSymbol:     "xxb",
			depositAmount:    sdk.NewDecCoin(common.NativeToken, sdk.ZeroInt()),
			totalValueLocked: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(0)),
			yieldedTokenInfos: YieldedTokenInfos{
				{
					RemainingAmount: sdk.NewDecCoinFromDec("wwb", sdk.NewDec(0)),
				},
			},
			totalAccumulatedRewards: sdk.SysCoins{},
			isFinished:              true,
		},
	}

	for _, test := range tests {
		pool := NewFarmPool(
			test.owner, test.name, sdk.NewDecCoinFromDec(test.lockedSymbol, sdk.ZeroDec()), test.depositAmount, test.totalValueLocked,
			test.yieldedTokenInfos, test.totalAccumulatedRewards,
		)
		require.Equal(t, test.isFinished, pool.Finished())
	}
}
