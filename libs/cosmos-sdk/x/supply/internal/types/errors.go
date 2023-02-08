package types

import (
	sdkerrors "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDeflation = sdkerrors.Register(ModuleName, 1, "failed. the deflation is larger than the current supply")
)
