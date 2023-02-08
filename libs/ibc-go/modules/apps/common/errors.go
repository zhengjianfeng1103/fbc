package common

import (
	"errors"
	"fmt"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types"

	"github.com/gogo/protobuf/proto"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	sdkerrors "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/errors"
)

// IBC port sentinel errors
var (
	ErrDisableProxyBeforeHeight = sdkerrors.Register(ModuleProxy, 1, "this feature is disable")
)

func MsgNotSupportBeforeHeight(msg proto.Message, h int64) error {
	if types.HigherThanVenus4(h) {
		return nil
	}
	return errors.New(fmt.Sprintf("msg:%s not support before height:%d", sdk.MsgTypeURL(msg), types.GetVenus4Height()))
}
