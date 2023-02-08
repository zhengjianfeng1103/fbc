package types

import "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
