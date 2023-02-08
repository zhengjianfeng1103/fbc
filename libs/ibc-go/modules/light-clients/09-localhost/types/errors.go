package types

import (
	sdkerrors "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/errors"
)

// Localhost sentinel errors
var (
	ErrConsensusStatesNotStored = sdkerrors.Register(SubModuleName, 2, "localhost does not store consensus states")
)
