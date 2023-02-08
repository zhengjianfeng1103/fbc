package auth

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/exported"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
