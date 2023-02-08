package types

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	authexported "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/exported"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/params"
	govtypes "github.com/zhengjianfeng1103/fbc/x/gov/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
}

// SupplyKeeper defines the expected interface needed to retrieve account balances.
type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type Subspace interface {
	GetParamSet(ctx sdk.Context, ps params.ParamSet)
	SetParamSet(ctx sdk.Context, ps params.ParamSet)
}

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}

type EvmKeeper interface {
	AddInnerTx(...interface{})
	DeleteInnerTx(...interface{})
}
