package auth

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	cliContext "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	codectypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	cosmost "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/store/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/module"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/upgrade"

	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/types"
	authinternaltypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/typesadapter"
)

var (
	_ module.AppModuleBasicAdapter = AppModuleBasic{}
	_ module.AppModuleAdapter      = AppModule{}
)

func (am AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (am AppModuleBasic) RegisterGRPCGatewayRoutes(clictx cliContext.CLIContext, serveMux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), serveMux, types.NewQueryClient(clictx))
}

func (am AppModuleBasic) RegisterRouterForGRPC(clictx cliContext.CLIContext, r *mux.Router) {
}

// ////
func (am AppModule) RegisterTask() upgrade.HeightTask {
	return nil
}

func (am AppModule) CommitFilter() *cosmost.StoreFilter {
	return nil
}

func (am AppModule) PruneFilter() *cosmost.StoreFilter {
	return nil
}

func (am AppModule) VersionFilter() *cosmost.VersionFilter {
	return nil
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	a := &am.accountKeeper
	types.RegisterQueryServer(cfg.QueryServer(), a)
}

func (am AppModuleBasic) GetTxCmdV2(cdc *codec.CodecProxy, reg codectypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func (AppModuleBasic) GetQueryCmdV2(cdc *codec.CodecProxy, reg codectypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*fmt.Stringer)(nil),
		&authinternaltypes.BaseAccount{},
	)
}
