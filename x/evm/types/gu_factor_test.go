package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

func TestMarshalGuFactor(t *testing.T) {
	str := "{\"gu_factor\":\"6000.000000000000000000\"}"
	factor, err := UnmarshalGuFactor(str)
	require.NoError(t, err)

	result := factor.Factor.MulInt(sdk.NewIntFromUint64(1220)).TruncateInt().Uint64()
	t.Log("result", result)

	t.Log("-1", sdk.NewDec(-1).String(), sdk.NewDec(-1).IsNegative())

}
