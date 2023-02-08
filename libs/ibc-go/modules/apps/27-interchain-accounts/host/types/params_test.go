package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/apps/27-interchain-accounts/host/types"
)

func TestValidateParams(t *testing.T) {
	require.NoError(t, types.DefaultParams().Validate())
	require.NoError(t, types.NewParams(false, []string{}).Validate())
}
