package simulation

import (
	"math/rand"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/simulation"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simulation.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
