package keeper_test

import (
	"fmt"

	"github.com/FiboChain/fbc/app/crypto/ethsecp256k1"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	"github.com/FiboChain/fbc/x/feesplit/types"
	"github.com/ethereum/go-ethereum/common"
)

func (suite *KeeperTestSuite) TestGetFees() {
	var expRes []types.FeeSplit

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"no fee splits registered",
			func() { expRes = []types.FeeSplit{} },
		},
		{
			"one fee split registered with withdraw address",
			func() {
				feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				expRes = []types.FeeSplit{feeSplit}
			},
		},
		{
			"one fee split registered with no withdraw address",
			func() {
				feeSplit := types.NewFeeSplit(contract, deployer, nil)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				expRes = []types.FeeSplit{feeSplit}
			},
		},
		{
			"multiple fee splits registered",
			func() {
				deployer2 := sdk.AccAddress(ethsecp256k1.GenerateAddress().Bytes())
				contract2 := ethsecp256k1.GenerateAddress()
				contract3 := ethsecp256k1.GenerateAddress()
				feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
				feeSplit2 := types.NewFeeSplit(contract2, deployer, nil)
				feeSplit3 := types.NewFeeSplit(contract3, deployer2, nil)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit2)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit3)
				expRes = []types.FeeSplit{feeSplit, feeSplit2, feeSplit3}
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			tc.malleate()

			res := suite.app.FeeSplitKeeper.GetFeeSplits(suite.ctx)
			suite.Require().ElementsMatch(expRes, res, tc.name)
		})
	}
}

func (suite *KeeperTestSuite) TestIterateFees() {
	var expRes []types.FeeSplit

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"no fee splits registered",
			func() { expRes = []types.FeeSplit{} },
		},
		{
			"one fee split registered with withdraw address",
			func() {
				feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				expRes = []types.FeeSplit{
					types.NewFeeSplit(contract, deployer, withdraw),
				}
			},
		},
		{
			"one fee split registered with no withdraw address",
			func() {
				feeSplit := types.NewFeeSplit(contract, deployer, nil)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				expRes = []types.FeeSplit{
					types.NewFeeSplit(contract, deployer, nil),
				}
			},
		},
		{
			"multiple fee splits registered",
			func() {
				deployer2 := sdk.AccAddress(ethsecp256k1.GenerateAddress().Bytes())
				contract2 := ethsecp256k1.GenerateAddress()
				contract3 := ethsecp256k1.GenerateAddress()
				feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
				feeSplit2 := types.NewFeeSplit(contract2, deployer, nil)
				feeSplit3 := types.NewFeeSplit(contract3, deployer2, nil)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit2)
				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit3)
				expRes = []types.FeeSplit{feeSplit, feeSplit2, feeSplit3}
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			tc.malleate()

			suite.app.FeeSplitKeeper.IterateFeeSplits(suite.ctx, func(feeSplit types.FeeSplit) (stop bool) {
				suite.Require().Contains(expRes, feeSplit, tc.name)
				return false
			})
		})
	}
}

func (suite *KeeperTestSuite) TestGetFeeSplit() {
	testCases := []struct {
		name        string
		contract    common.Address
		deployer    sdk.AccAddress
		withdraw    sdk.AccAddress
		found       bool
		expWithdraw bool
	}{
		{
			"fee with no withdraw address",
			contract,
			deployer,
			nil,
			true,
			false,
		},
		{
			"fee with withdraw address same as deployer",
			contract,
			deployer,
			deployer,
			true,
			false,
		},
		{
			"fee with withdraw address same as contract",
			contract,
			deployer,
			sdk.AccAddress(contract.Bytes()),
			true,
			true,
		},
		{
			"fee with withdraw address different than deployer",
			contract,
			deployer,
			withdraw,
			true,
			true,
		},
		{
			"no fee",
			common.Address{},
			nil,
			nil,
			false,
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			if tc.found {
				feeSplit := types.NewFeeSplit(tc.contract, tc.deployer, tc.withdraw)
				if tc.deployer.Equals(tc.withdraw) {
					feeSplit.WithdrawerAddress = nil
				}

				suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
				suite.app.FeeSplitKeeper.SetDeployerMap(suite.ctx, tc.deployer, tc.contract)
			}

			if tc.expWithdraw {
				suite.app.FeeSplitKeeper.SetWithdrawerMap(suite.ctx, tc.withdraw, tc.contract)
			}

			feeSplit, found := suite.app.FeeSplitKeeper.GetFeeSplit(suite.ctx, tc.contract)
			foundD := suite.app.FeeSplitKeeper.IsDeployerMapSet(suite.ctx, tc.deployer, tc.contract)
			foundW := suite.app.FeeSplitKeeper.IsWithdrawerMapSet(suite.ctx, tc.withdraw, tc.contract)

			if tc.found {
				suite.Require().True(found, tc.name)
				suite.Require().Equal(tc.deployer, feeSplit.DeployerAddress, tc.name)
				suite.Require().Equal(tc.contract, feeSplit.ContractAddress, tc.name)

				suite.Require().True(foundD, tc.name)

				if tc.expWithdraw {
					suite.Require().Equal(tc.withdraw, feeSplit.WithdrawerAddress, tc.name)
					suite.Require().True(foundW, tc.name)
				} else {
					suite.Require().Equal(tc.deployer, feeSplit.WithdrawerAddress, tc.name)
					suite.Require().False(foundW, tc.name)
				}
			} else {
				suite.Require().False(found, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteFeeSplit() {
	feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
	suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)

	initialFee, found := suite.app.FeeSplitKeeper.GetFeeSplit(suite.ctx, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing fee split", func() {}, true},
		{
			"deleted fee split",
			func() {
				suite.app.FeeSplitKeeper.DeleteFeeSplit(suite.ctx, feeSplit)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		feeSplit, found := suite.app.FeeSplitKeeper.GetFeeSplit(suite.ctx, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
			suite.Require().Equal(initialFee, feeSplit, tc.name)
		} else {
			suite.Require().False(found, tc.name)
			suite.Require().Equal(types.FeeSplit{}, feeSplit, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestDeleteDeployerMap() {
	suite.app.FeeSplitKeeper.SetDeployerMap(suite.ctx, deployer, contract)
	found := suite.app.FeeSplitKeeper.IsDeployerMapSet(suite.ctx, deployer, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing deployer", func() {}, true},
		{
			"deleted deployer",
			func() {
				suite.app.FeeSplitKeeper.DeleteDeployerMap(suite.ctx, deployer, contract)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		found := suite.app.FeeSplitKeeper.IsDeployerMapSet(suite.ctx, deployer, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestDeleteWithdrawMap() {
	suite.app.FeeSplitKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)
	found := suite.app.FeeSplitKeeper.IsWithdrawerMapSet(suite.ctx, withdraw, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing withdraw", func() {}, true},
		{
			"deleted withdraw",
			func() {
				suite.app.FeeSplitKeeper.DeleteWithdrawerMap(suite.ctx, withdraw, contract)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		found := suite.app.FeeSplitKeeper.IsWithdrawerMapSet(suite.ctx, withdraw, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestIsFeeSplitRegistered() {
	feeSplit := types.NewFeeSplit(contract, deployer, withdraw)
	suite.app.FeeSplitKeeper.SetFeeSplit(suite.ctx, feeSplit)
	_, found := suite.app.FeeSplitKeeper.GetFeeSplit(suite.ctx, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		contract common.Address
		ok       bool
	}{
		{"registered fee split", contract, true},
		{"fee split not registered", common.Address{}, false},
		{"fee split not registered", ethsecp256k1.GenerateAddress(), false},
	}
	for _, tc := range testCases {
		found := suite.app.FeeSplitKeeper.IsFeeSplitRegistered(suite.ctx, tc.contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}
