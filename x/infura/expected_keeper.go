package infura

import evm "github.com/zhengjianfeng1103/fbc/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
