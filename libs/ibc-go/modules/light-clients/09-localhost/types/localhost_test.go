package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/zhengjianfeng1103/fbc/libs/tendermint/abci/types"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	clienttypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/exported"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/testing/simapp"
)

const (
	height = 4
)

var (
	clientHeight = clienttypes.NewHeight(0, 10)
)

type LocalhostTestSuite struct {
	suite.Suite

	cdc   *codec.CodecProxy
	ctx   sdk.Context
	store sdk.KVStore
}

func (suite *LocalhostTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(isCheckTx)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{Height: 1, ChainID: "ibc-chain"})
	suite.store = app.IBCKeeper.V2Keeper.ClientKeeper.ClientStore(suite.ctx, exported.Localhost)
}

func TestLocalhostTestSuite(t *testing.T) {
	suite.Run(t, new(LocalhostTestSuite))
}
