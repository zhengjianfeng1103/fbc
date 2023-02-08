package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store/mpt"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

type KeeperMptTestSuite struct {
	KeeperTestSuite
}

func (suite *KeeperMptTestSuite) SetupTest() {
	mpt.TrieWriteAhead = true
	types.UnittestOnlySetMilestoneMarsHeight(1)

	suite.KeeperTestSuite.SetupTest()
}

func TestKeeperMptTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperMptTestSuite))
}
