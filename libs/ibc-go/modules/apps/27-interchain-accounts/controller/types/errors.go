package types

import (
	sdkerrors "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/errors"
)

// ICA Controller sentinel errors
var (
	ErrControllerSubModuleDisabled = sdkerrors.Register(SubModuleName, 2, "controller submodule is disabled")
)
