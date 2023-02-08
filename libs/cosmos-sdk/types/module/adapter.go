package module

import (
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	clictx "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	codectypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

// AppModuleBasic is the standard form for basic non-dependant elements of an application module.
type AppModuleBasicAdapter interface {
	AppModuleBasic
	RegisterInterfaces(codectypes.InterfaceRegistry)
	// client functionality
	RegisterGRPCGatewayRoutes(clictx.CLIContext, *runtime.ServeMux)
	GetTxCmdV2(cdc *codec.CodecProxy, reg codectypes.InterfaceRegistry) *cobra.Command
	GetQueryCmdV2(cdc *codec.CodecProxy, reg codectypes.InterfaceRegistry) *cobra.Command

	RegisterRouterForGRPC(cliCtx clictx.CLIContext, r *mux.Router)
}

// AppModuleGenesis is the standard form for an application module genesis functions
type AppModuleGenesisAdapter interface {
	AppModuleGenesis
	AppModuleBasicAdapter
}

// AppModule is the standard form for an application module
type AppModuleAdapter interface {
	AppModule
	AppModuleGenesisAdapter
	// registers
	RegisterInvariants(sdk.InvariantRegistry)
	// RegisterServices allows a module to register services
	RegisterServices(Configurator)
}
