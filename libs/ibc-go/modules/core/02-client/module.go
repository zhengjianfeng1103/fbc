package client

import (
	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	interfacetypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/client/cli"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}

// GetQueryCmd returns no root query command for the IBC client
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	return cli.GetQueryCmd(cdc, reg)
}

// GetTxCmd returns the root tx command for 02-client.
func GetTxCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	return cli.NewTxCmd(cdc, reg)
}

// RegisterQueryService registers the gRPC query service for IBC client.
func RegisterQueryService(server grpc.Server, queryServer types.QueryServer) {
	types.RegisterQueryServer(server, queryServer)
}
