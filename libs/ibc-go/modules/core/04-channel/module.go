package channel

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	interfacetypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/04-channel/client/cli"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/04-channel/types"
)

// Name returns the IBC channel ICS name.
func Name() string {
	return types.SubModuleName
}

// GetTxCmd returns the root tx command for IBC channels.
func GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for IBC channels.
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	return cli.GetQueryCmd(cdc, reg)
}

// RegisterQueryService registers the gRPC query service for IBC channels.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
