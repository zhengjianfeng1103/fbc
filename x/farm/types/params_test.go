//go:build ignore
// +build ignore

package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

const (
	strExpected = `Params:
  Quote Symbol:								usdf
  Create Pool Fee:							0.000000000000000000` + sdk.DefaultBondDenom + `
  Create Pool Deposit:						10.000000000000000000` + sdk.DefaultBondDenom + `
  Yield Native Token Enabled:               false`
)

func TestParams(t *testing.T) {
	defaultState := DefaultGenesisState()
	defaultParams := DefaultParams()

	require.Equal(t, defaultState.Params, defaultParams)
	require.Equal(t, strExpected, defaultParams.String())
}
