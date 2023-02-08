// nolint
// ALIASGEN: github.com/zhengjianfeng1103/fbc/x/staking/types
package staking

import (
	"github.com/zhengjianfeng1103/fbc/x/staking/types"
)

var (
	// functions aliases
	NewCommissionRates                = types.NewCommissionRates
	NewMsgEditValidatorCommissionRate = types.NewMsgEditValidatorCommissionRate
	NewMsgDestroyValidator            = types.NewMsgDestroyValidator
	NewMsgRegProxy                    = types.NewMsgRegProxy
	NewMsgBindProxy                   = types.NewMsgBindProxy
	NewMsgUnbindProxy                 = types.NewMsgUnbindProxy
)
