package types

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	codectypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	sdkerrors "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/errors"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/tx/signing"
	clienttypes "github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/exported"
)

// RegisterInterfaces register the ibc channel submodule interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*exported.Misbehaviour)(nil),
		&Misbehaviour{},
	)
}

func UnmarshalSignatureData(cdc *codec.CodecProxy, data []byte) (signing.SignatureData, error) {
	protoSigData := &signing.SignatureDescriptor_Data{}
	if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, protoSigData); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to unmarshal proof into type %T", protoSigData)
	}

	sigData := signing.SignatureDataFromProto(protoSigData)

	return sigData, nil
}

// UnmarshalDataByType attempts to unmarshal the data to the specified type. An error is
// return if it fails.
func UnmarshalDataByType(cdc *codec.CodecProxy, dataType DataType, data []byte) (Data, error) {
	if len(data) == 0 {
		return nil, sdkerrors.Wrap(ErrInvalidSignatureAndData, "data cannot be empty")
	}

	switch dataType {
	case UNSPECIFIED:
		return nil, sdkerrors.Wrap(ErrInvalidDataType, "data type cannot be UNSPECIFIED")

	case CLIENT:
		clientData := &ClientStateData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, clientData); err != nil {
			return nil, err
		}

		// unpack any
		if _, err := clienttypes.UnpackClientState(clientData.ClientState); err != nil {
			return nil, err
		}
		return clientData, nil

	case CONSENSUS:
		consensusData := &ConsensusStateData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, consensusData); err != nil {
			return nil, err
		}

		// unpack any
		if _, err := clienttypes.UnpackConsensusState(consensusData.ConsensusState); err != nil {
			return nil, err
		}
		return consensusData, nil

	case CONNECTION:
		connectionData := &ConnectionStateData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, connectionData); err != nil {
			return nil, err
		}

		return connectionData, nil

	case CHANNEL:
		channelData := &ChannelStateData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, channelData); err != nil {
			return nil, err
		}

		return channelData, nil

	case PACKETCOMMITMENT:
		commitmentData := &PacketCommitmentData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, commitmentData); err != nil {
			return nil, err
		}

		return commitmentData, nil

	case PACKETACKNOWLEDGEMENT:
		ackData := &PacketAcknowledgementData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, ackData); err != nil {
			return nil, err
		}

		return ackData, nil

	case PACKETRECEIPTABSENCE:
		receiptAbsenceData := &PacketReceiptAbsenceData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, receiptAbsenceData); err != nil {
			return nil, err
		}

		return receiptAbsenceData, nil

	case NEXTSEQUENCERECV:
		nextSeqRecvData := &NextSequenceRecvData{}
		if err := cdc.GetProtocMarshal().UnmarshalBinaryBare(data, nextSeqRecvData); err != nil {
			return nil, err
		}

		return nextSeqRecvData, nil

	default:
		return nil, sdkerrors.Wrapf(ErrInvalidDataType, "unsupported data type %T", dataType)
	}
}
