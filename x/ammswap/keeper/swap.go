package keeper

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/x/ammswap/types"
)

// IsTokenExist check token is exist
func (k Keeper) IsTokenExist(ctx sdk.Context, token string) error {
	isExist := k.tokenKeeper.TokenExist(ctx, token)
	if !isExist {
		return types.ErrTokenNotExist()
	}

	t := k.tokenKeeper.GetTokenInfo(ctx, token)
	if t.Type == types.GenerateTokenType {
		return types.ErrInvalidCoins()
	}
	return nil

}
