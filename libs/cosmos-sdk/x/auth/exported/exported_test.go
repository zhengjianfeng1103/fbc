package exported_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/crypto/secp256k1"

	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/exported"
	authtypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/types"
)

func TestGenesisAccountsContains(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	acc := authtypes.NewBaseAccount(addr, nil, secp256k1.GenPrivKey().PubKey(), 0, 0)

	genAccounts := exported.GenesisAccounts{}
	require.False(t, genAccounts.Contains(acc.GetAddress()))

	genAccounts = append(genAccounts, acc)
	require.True(t, genAccounts.Contains(acc.GetAddress()))
}
