package simulator

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

type Simulator interface {
	Simulate([]sdk.Msg) (*sdk.Result, error)
	Context() *sdk.Context
}

var NewWasmSimulator func() Simulator
