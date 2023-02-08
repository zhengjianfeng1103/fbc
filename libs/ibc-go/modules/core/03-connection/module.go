package connection

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	interfacetypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/03-connection/client/cli"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/03-connection/types"
)

// Name returns the IBC connection ICS name.
func Name() string {
	return types.SubModuleName
}

// GetQueryCmd returns the root query command for the IBC connections.
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	return cli.GetQueryCmd(cdc, reg)
}

// RegisterQueryService registers the gRPC query service for IBC connections.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
