package types

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// GetDelegatorAddress gets delegator address
func (d Delegator) GetDelegatorAddress() sdk.AccAddress {
	return d.DelegatorAddress
}
