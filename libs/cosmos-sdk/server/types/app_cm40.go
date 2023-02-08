package app

import (
	gogogrpc "github.com/gogo/protobuf/grpc"
	cliContext "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
)

type ApplicationAdapter interface {
	RegisterGRPCServer(gogogrpc.Server)
	RegisterTxService(clientCtx cliContext.CLIContext)
}
