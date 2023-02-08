package main

import (
	"fmt"

	"github.com/zhengjianfeng1103/fbc/app"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/tx"
	"github.com/zhengjianfeng1103/fbc/x/wasm/proxy"

	mintclient "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint/client"
	mintrest "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint/client/rest"
	erc20client "github.com/zhengjianfeng1103/fbc/x/erc20/client"
	erc20rest "github.com/zhengjianfeng1103/fbc/x/erc20/client/rest"
	evmclient "github.com/zhengjianfeng1103/fbc/x/evm/client"

	"github.com/spf13/viper"
	"github.com/zhengjianfeng1103/fbc/app/rpc"
	"github.com/zhengjianfeng1103/fbc/app/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/lcd"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/server"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth"
	authrest "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/bank/client/rest"
	supplyrest "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/supply/client/rest"
	ammswaprest "github.com/zhengjianfeng1103/fbc/x/ammswap/client/rest"
	dexclient "github.com/zhengjianfeng1103/fbc/x/dex/client"
	dexrest "github.com/zhengjianfeng1103/fbc/x/dex/client/rest"
	dist "github.com/zhengjianfeng1103/fbc/x/distribution"
	distr "github.com/zhengjianfeng1103/fbc/x/distribution"
	distrest "github.com/zhengjianfeng1103/fbc/x/distribution/client/rest"
	evmrest "github.com/zhengjianfeng1103/fbc/x/evm/client/rest"
	farmclient "github.com/zhengjianfeng1103/fbc/x/farm/client"
	farmrest "github.com/zhengjianfeng1103/fbc/x/farm/client/rest"
	fsrest "github.com/zhengjianfeng1103/fbc/x/feesplit/client/rest"
	govrest "github.com/zhengjianfeng1103/fbc/x/gov/client/rest"
	orderrest "github.com/zhengjianfeng1103/fbc/x/order/client/rest"
	paramsclient "github.com/zhengjianfeng1103/fbc/x/params/client"
	stakingrest "github.com/zhengjianfeng1103/fbc/x/staking/client/rest"
	"github.com/zhengjianfeng1103/fbc/x/token"
	tokensrest "github.com/zhengjianfeng1103/fbc/x/token/client/rest"
	wasmrest "github.com/zhengjianfeng1103/fbc/x/wasm/client/rest"
)

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerGrpc(rs)
	rpc.RegisterRoutes(rs)
	pathPrefix := viper.GetString(server.FlagRestPathPrefix)
	if pathPrefix == "" {
		pathPrefix = types.EthBech32Prefix
	}
	registerRoutesV1(rs, pathPrefix)
	registerRoutesV2(rs, pathPrefix)
	proxy.SetCliContext(rs.CliCtx)
}

func registerGrpc(rs *lcd.RestServer) {
	app.ModuleBasics.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
	app.ModuleBasics.RegisterRPCRouterForGRPC(rs.CliCtx, rs.Mux)
	tx.RegisterGRPCGatewayRoutes(rs.CliCtx, rs.GRPCGatewayRouter)
}

func registerRoutesV1(rs *lcd.RestServer, pathPrefix string) {
	v1Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v1", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v1Router)
	authrest.RegisterRoutes(rs.CliCtx, v1Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v1Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v1Router)
	distrest.RegisterRoutes(rs.CliCtx, v1Router, dist.StoreKey)

	orderrest.RegisterRoutes(rs.CliCtx, v1Router)
	tokensrest.RegisterRoutes(rs.CliCtx, v1Router, token.StoreKey)
	dexrest.RegisterRoutes(rs.CliCtx, v1Router)
	ammswaprest.RegisterRoutes(rs.CliCtx, v1Router)
	supplyrest.RegisterRoutes(rs.CliCtx, v1Router)
	farmrest.RegisterRoutes(rs.CliCtx, v1Router)
	evmrest.RegisterRoutes(rs.CliCtx, v1Router)
	erc20rest.RegisterRoutes(rs.CliCtx, v1Router)
	wasmrest.RegisterRoutes(rs.CliCtx, v1Router)
	fsrest.RegisterRoutes(rs.CliCtx, v1Router)
	govrest.RegisterRoutes(rs.CliCtx, v1Router,
		[]govrest.ProposalRESTHandler{
			paramsclient.ProposalHandler.RESTHandler(rs.CliCtx),
			distr.CommunityPoolSpendProposalHandler.RESTHandler(rs.CliCtx),
			distr.ChangeDistributionTypeProposalHandler.RESTHandler(rs.CliCtx),
			distr.WithdrawRewardEnabledProposalHandler.RESTHandler(rs.CliCtx),
			distr.RewardTruncatePrecisionProposalHandler.RESTHandler(rs.CliCtx),
			dexclient.DelistProposalHandler.RESTHandler(rs.CliCtx),
			farmclient.ManageWhiteListProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageContractDeploymentWhitelistProposalHandler.RESTHandler(rs.CliCtx),
			evmclient.ManageSysContractAddressProposalHandler.RESTHandler(rs.CliCtx),
			mintclient.ManageTreasuresProposalHandler.RESTHandler(rs.CliCtx),
			erc20client.TokenMappingProposalHandler.RESTHandler(rs.CliCtx),
		},
	)
	mintrest.RegisterRoutes(rs.CliCtx, v1Router)

}

func registerRoutesV2(rs *lcd.RestServer, pathPrefix string) {
	v2Router := rs.Mux.PathPrefix(fmt.Sprintf("/%s/v2", pathPrefix)).Name("v1").Subrouter()
	client.RegisterRoutes(rs.CliCtx, v2Router)
	authrest.RegisterRoutes(rs.CliCtx, v2Router, auth.StoreKey)
	bankrest.RegisterRoutes(rs.CliCtx, v2Router)
	stakingrest.RegisterRoutes(rs.CliCtx, v2Router)
	distrest.RegisterRoutes(rs.CliCtx, v2Router, dist.StoreKey)
	orderrest.RegisterRoutesV2(rs.CliCtx, v2Router)
	tokensrest.RegisterRoutesV2(rs.CliCtx, v2Router, token.StoreKey)
	fsrest.RegisterRoutesV2(rs.CliCtx, v2Router)
}
