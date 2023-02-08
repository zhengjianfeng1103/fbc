package keeper

import (
	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"

	//"github.com/stretchr/testify/require"
	"testing"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types/time"
)

func TestDecay(t *testing.T) {
	now := time.Now().Unix()
	after := time.Now().AddDate(0, 0, 52*7).Unix()

	tokens := sdk.NewDec(1000)
	nowDec, err := calculateWeight(now, tokens)
	require.NoError(t, err)
	afterDec, err := calculateWeight(after, tokens)
	require.NoError(t, err)
	require.Equal(t, sdk.NewDec(2), afterDec.Quo(nowDec))
}
