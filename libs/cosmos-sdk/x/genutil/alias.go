// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/genutil/types
package genutil

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/genutil/types"
)

const (
	ModuleName = types.ModuleName
)

var (
	// functions aliases
	NewGenesisState             = types.NewGenesisState
	NewGenesisStateFromStdTx    = types.NewGenesisStateFromStdTx
	NewInitConfig               = types.NewInitConfig
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	SetGenesisStateInAppState   = types.SetGenesisStateInAppState
	GenesisStateFromGenDoc      = types.GenesisStateFromGenDoc
	GenesisStateFromGenFile     = types.GenesisStateFromGenFile
	ValidateGenesis             = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	GenesisState      = types.GenesisState
	AppMap            = types.AppMap
	MigrationCallback = types.MigrationCallback
	MigrationMap      = types.MigrationMap
	InitConfig        = types.InitConfig
)
