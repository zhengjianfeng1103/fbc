package proto

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	abci "github.com/zhengjianfeng1103/fbc/libs/tendermint/abci/types"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/libs/log"
	dbm "github.com/zhengjianfeng1103/fbc/libs/tm-db"
)

func createTestInput(t *testing.T) (sdk.Context, ProtocolKeeper) {
	keyMain := sdk.NewKVStoreKey("main")

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyMain, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewTMLogger(os.Stdout))

	keeper := NewProtocolKeeper(keyMain)

	return ctx, keeper
}
