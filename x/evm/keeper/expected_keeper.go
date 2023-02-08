package keeper

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	govtypes "github.com/zhengjianfeng1103/fbc/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
