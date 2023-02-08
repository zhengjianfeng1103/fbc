package keeper

import (
	govtypes "github.com/zhengjianfeng1103/fbc/x/gov/types"

	"time"

	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time)
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
