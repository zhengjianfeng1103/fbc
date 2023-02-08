package ante

import (
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	clienttypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/04-channel/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/keeper"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

type AnteDecorator struct {
	k keeper.IBCServerKeeper
}

func NewAnteDecorator(k keeper.IBCServerKeeper) AnteDecorator {
	return AnteDecorator{k: k}
}

// AnteDecorator returns an error if a multiMsg tx only contains packet messages (Recv, Ack, Timeout) and additional update messages and all packet messages
// are redundant. If the transaction is just a single UpdateClient message, or the multimsg transaction contains some other message type, then the antedecorator returns no error
// and continues processing to ensure these transactions are included.
// This will ensure that relayers do not waste fees on multiMsg transactions when another relayer has already submitted all packets, by rejecting the tx at the mempool layer.
func (ad AnteDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !types.HigherThanVenus1(ctx.BlockHeight()) {
		return next(ctx, tx, simulate)
	}
	if types.HigherThanVenus4(ctx.BlockHeight()) {
		return ad.v4(ctx, tx, simulate, next)
	}
	// do not run redundancy check on DeliverTx or simulate
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {
		// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`
		redundancies := 0
		packetMsgs := 0
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *channeltypes.MsgRecvPacket:

				if _, found := ad.k.GetPacketReceipt(ctx, msg.Packet.GetDestPort(), msg.Packet.GetDestChannel(), msg.Packet.GetSequence()); found {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgAcknowledgement:
				if commitment := ad.k.GetPacketCommitment(ctx, msg.Packet.GetSourcePort(), msg.Packet.GetSourceChannel(), msg.Packet.GetSequence()); len(commitment) == 0 {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeout:
				if commitment := ad.k.GetPacketCommitment(ctx, msg.Packet.GetSourcePort(), msg.Packet.GetSourceChannel(), msg.Packet.GetSequence()); len(commitment) == 0 {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeoutOnClose:
				if commitment := ad.k.GetPacketCommitment(ctx, msg.Packet.GetSourcePort(), msg.Packet.GetSourceChannel(), msg.Packet.GetSequence()); len(commitment) == 0 {
					redundancies += 1
				}
				packetMsgs += 1

			case *clienttypes.MsgUpdateClient:
				// do nothing here, as we want to avoid updating clients if it is batched with only redundant messages

			default:
				// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
				// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
				// even if they get batched with redundant packet messages.
				return next(ctx, tx, simulate)
			}

		}

		// only return error if all packet messages are redundant
		if redundancies == packetMsgs && packetMsgs > 0 {
			return ctx, channeltypes.ErrRedundantTx
		}
	}
	return next(ctx, tx, simulate)
}

func (ad AnteDecorator) v4(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// do not run redundancy check on DeliverTx or simulate
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {
		// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`
		redundancies := 0
		packetMsgs := 0
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *channeltypes.MsgRecvPacket:
				response, err := ad.k.RecvPacket(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgAcknowledgement:
				response, err := ad.k.Acknowledgement(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeout:
				response, err := ad.k.Timeout(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *channeltypes.MsgTimeoutOnClose:
				response, err := ad.k.TimeoutOnClose(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies += 1
				}
				packetMsgs += 1

			case *clienttypes.MsgUpdateClient:
				_, err := ad.k.UpdateClient(sdk.WrapSDKContext(ctx), msg)
				if err != nil {
					return ctx, err
				}

			default:
				// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
				// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
				// even if they get batched with redundant packet messages.
				return next(ctx, tx, simulate)
			}
		}

		// only return error if all packet messages are redundant
		if redundancies == packetMsgs && packetMsgs > 0 {
			return ctx, channeltypes.ErrRedundantTx
		}
	}
	return next(ctx, tx, simulate)
}
