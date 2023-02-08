package store

import (
	dbm "github.com/zhengjianfeng1103/fbc/libs/tm-db"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store/cache"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store/rootmulti"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
