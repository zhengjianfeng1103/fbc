package appstatus

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/spf13/viper"
	bam "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/baseapp"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/flags"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth"
	capabilitytypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/capability/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/params"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/supply"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/upgrade"
	"github.com/zhengjianfeng1103/fbc/libs/iavl"
	ibctransfertypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/apps/transfer/types"
	ibchost "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/24-host"
	dbm "github.com/zhengjianfeng1103/fbc/libs/tm-db"
	"github.com/zhengjianfeng1103/fbc/x/ammswap"
	dex "github.com/zhengjianfeng1103/fbc/x/dex/types"
	distr "github.com/zhengjianfeng1103/fbc/x/distribution"
	"github.com/zhengjianfeng1103/fbc/x/erc20"
	"github.com/zhengjianfeng1103/fbc/x/evidence"
	"github.com/zhengjianfeng1103/fbc/x/evm"
	"github.com/zhengjianfeng1103/fbc/x/farm"
	"github.com/zhengjianfeng1103/fbc/x/feesplit"
	"github.com/zhengjianfeng1103/fbc/x/gov"
	"github.com/zhengjianfeng1103/fbc/x/order"
	"github.com/zhengjianfeng1103/fbc/x/slashing"
	staking "github.com/zhengjianfeng1103/fbc/x/staking/types"
	token "github.com/zhengjianfeng1103/fbc/x/token/types"
)

const (
	applicationDB = "application"
	dbFolder      = "data"
)

func GetAllStoreKeys() []string {
	return []string{
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,
		evm.StoreKey, token.StoreKey, token.KeyLock, dex.StoreKey, dex.TokenPairStoreKey,
		order.OrderStoreKey, ammswap.StoreKey, farm.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		ibchost.StoreKey,
		erc20.StoreKey,
		// mpt.StoreKey,
		// wasm.StoreKey,
		feesplit.StoreKey,
	}
}

func IsFastStorageStrategy() bool {
	return checkFastStorageStrategy(GetAllStoreKeys())
}

func checkFastStorageStrategy(storeKeys []string) bool {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, v := range storeKeys {
		if !isFss(db, v) {
			return false
		}
	}

	return true
}

func isFss(db dbm.DB, storeKey string) bool {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	return iavl.IsFastStorageStrategy(prefixDB)
}

func GetFastStorageVersion() int64 {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	storeKeys := GetAllStoreKeys()
	var ret int64 = math.MaxInt64
	for _, v := range storeKeys {
		version := getVersion(db, v)
		if version < ret {
			ret = version
		}
	}

	return ret
}

func getVersion(db dbm.DB, storeKey string) int64 {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	version, _ := iavl.GetFastStorageVersion(prefixDB)

	return version
}
