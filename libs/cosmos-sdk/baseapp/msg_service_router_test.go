package baseapp_test

import (
	"os"
	"testing"

	fbexchaincodec "github.com/zhengjianfeng1103/fbc/app/codec"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/testing/simapp"
	"github.com/zhengjianfeng1103/fbc/x/evm"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/libs/log"

	"github.com/stretchr/testify/require"
	dbm "github.com/zhengjianfeng1103/fbc/libs/tm-db"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/baseapp"

	"github.com/zhengjianfeng1103/fbc/x/evm/types/testdata"
)

func TestRegisterMsgService(t *testing.T) {
	db := dbm.NewMemDB()

	// Create an encoding config that doesn't register testdata Msg services.
	codecProxy, interfaceRegistry := fbexchaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Register testdata Msg services, and rerun `RegisterService`.
	testdata.RegisterInterfaces(interfaceRegistry)
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}

func TestRegisterMsgServiceTwice(t *testing.T) {
	// Setup baseapp.
	db := dbm.NewMemDB()
	codecProxy, interfaceRegistry := fbexchaincodec.MakeCodecSuit(simapp.ModuleBasics)
	app := baseapp.NewBaseApp("test", log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, evm.TxDecoder(codecProxy))
	app.SetInterfaceRegistry(interfaceRegistry)
	testdata.RegisterInterfaces(interfaceRegistry)

	// First time registering service shouldn't panic.
	require.NotPanics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})

	// Second time should panic.
	require.Panics(t, func() {
		testdata.RegisterMsgServer(
			app.MsgServiceRouter(),
			testdata.MsgServerImpl{},
		)
	})
}
