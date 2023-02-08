package types

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// SupplyKeeper defines the expected supply keeper (noalias)
type SupplyKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
