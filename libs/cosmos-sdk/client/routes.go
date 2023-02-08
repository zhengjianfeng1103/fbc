package client

import (
	"github.com/gorilla/mux"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
