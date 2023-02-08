package vmbridge

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/module"
	"github.com/zhengjianfeng1103/fbc/x/vmbridge/keeper"
	"github.com/zhengjianfeng1103/fbc/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
