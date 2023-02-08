package evidence

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/zhengjianfeng1103/fbc/libs/tendermint/crypto/encoding/amino"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
	cryptoamino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
