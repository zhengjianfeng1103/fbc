package auth

import (
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/exported"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
