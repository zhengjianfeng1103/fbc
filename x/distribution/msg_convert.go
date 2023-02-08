package distribution

import (
	"encoding/json"
	"errors"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/baseapp"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	tmtypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
	"github.com/zhengjianfeng1103/fbc/x/common"
	"github.com/zhengjianfeng1103/fbc/x/distribution/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	enableHeight := tmtypes.GetVenus3Height()
	baseapp.RegisterCmHandle("fbexchain/distribution/MsgWithdrawDelegatorAllRewards", baseapp.NewCMHandle(ConvertWithdrawDelegatorAllRewardsMsg, enableHeight))
}

func ConvertWithdrawDelegatorAllRewardsMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdrawDelegatorAllRewards{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
