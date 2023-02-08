package node

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/zhengjianfeng1103/fbc/libs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
