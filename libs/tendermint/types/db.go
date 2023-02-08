package types

import dbm "github.com/zhengjianfeng1103/fbc/libs/tm-db"

// DBBackend This is set at compile time.
var DBBackend = string(dbm.GoLevelDBBackend)
