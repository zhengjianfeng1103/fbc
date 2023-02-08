//go:build ignore
// +build ignore

package keeper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/x/farm/types"
	govtypes "github.com/zhengjianfeng1103/fbc/x/gov/types"
)

func TestCheckMsgSubmitProposal(t *testing.T) {
	ctx, k := GetKeeper(t)

	proposal := govtypes.Proposal{Content: types.NewManageWhiteListProposal(
		"Test",
		"description",
		"pool",
		true,
	)}

	require.Equal(t, sdk.SysCoins(nil), k.GetMinDeposit(ctx, MockContent{}))
	require.Equal(t, k.govKeeper.GetDepositParams(ctx).MinDeposit, k.GetMinDeposit(ctx, proposal.Content))

	require.Equal(t, time.Duration(0), k.GetMaxDepositPeriod(ctx, MockContent{}))
	require.Equal(t, k.govKeeper.GetDepositParams(ctx).MaxDepositPeriod, k.GetMaxDepositPeriod(ctx, proposal.Content))

	require.Equal(t, time.Duration(0), k.GetVotingPeriod(ctx, MockContent{}))
	require.Equal(t, k.govKeeper.GetVotingParams(ctx).VotingPeriod, k.GetVotingPeriod(ctx, proposal.Content))

	require.Error(t, k.CheckMsgSubmitProposal(ctx, govtypes.MsgSubmitProposal{Content: MockContent{}}))
	err := k.CheckMsgSubmitProposal(ctx, govtypes.MsgSubmitProposal{Content: proposal.Content})
	require.Error(t, err)
}

func TestCheckMsgManageWhiteListProposal(t *testing.T) {
	ctx, k := GetKeeper(t)
	quoteSymbol := types.DefaultParams().QuoteSymbol

	proposal := types.NewManageWhiteListProposal(
		"Test",
		"description",
		"pool",
		false,
	)

	err := k.CheckMsgManageWhiteListProposal(ctx, proposal)
	require.Error(t, err)

	k.SetWhitelist(ctx, proposal.PoolName)
	err = k.CheckMsgManageWhiteListProposal(ctx, proposal)
	require.NoError(t, err)

	proposal.IsAdded = true
	err = k.CheckMsgManageWhiteListProposal(ctx, proposal)
	require.Error(t, err)

	lockedSymbol := "xxb"
	pool := types.FarmPool{
		Name:          proposal.PoolName,
		MinLockAmount: sdk.NewDecCoinFromDec(lockedSymbol, sdk.ZeroDec()),
	}
	k.SetFarmPool(ctx, pool)
	err = k.CheckMsgManageWhiteListProposal(ctx, proposal)
	require.Error(t, err)

	SetSwapTokenPair(ctx, k.Keeper, lockedSymbol, quoteSymbol)
	err = k.CheckMsgManageWhiteListProposal(ctx, proposal)
	require.NoError(t, err)
}
