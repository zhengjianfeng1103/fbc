package params

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/x/params/types"
)

// ModuleCdc is the codec of module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterCodec registers all necessary param module types with a given codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.ParameterChangeProposal{}, "fbexchain/params/ParameterChangeProposal", nil)
}
