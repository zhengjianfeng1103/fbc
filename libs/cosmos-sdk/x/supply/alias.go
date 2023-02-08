// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/supply/internal/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/supply/internal/types
package supply

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/supply/internal/keeper"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/supply/internal/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
	Minter       = types.Minter
	Burner       = types.Burner
	Staking      = types.Staking
)

var (
	// functions aliases
	RegisterInvariants     = keeper.RegisterInvariants
	AllInvariants          = keeper.AllInvariants
	TotalSupply            = keeper.TotalSupply
	NewKeeper              = keeper.NewKeeper
	NewSupplyKeeperAdapter = keeper.NewSupplyKeerAdapter
	NewQuerier             = keeper.NewQuerier
	SupplyKey              = keeper.SupplyKey
	NewModuleAddress       = types.NewModuleAddress
	NewEmptyModuleAccount  = types.NewEmptyModuleAccount
	NewModuleAccount       = types.NewModuleAccount
	RegisterCodec          = types.RegisterCodec
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	NewSupply              = types.NewSupply
	DefaultSupply          = types.DefaultSupply

	NewQueryTotalSupplyParams = types.NewQueryTotalSupplyParams
	QueryTotalSupply          = types.QueryTotalSupply

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper        = keeper.Keeper
	KeeperAdapter = keeper.SupplyKeerAdapter
	ModuleAccount = types.ModuleAccount
	GenesisState  = types.GenesisState
	Supply        = types.Supply
)
