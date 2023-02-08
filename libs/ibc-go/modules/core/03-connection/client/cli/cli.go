package cli

import (
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	interfacetypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/03-connection/types"
)

// GetQueryCmd returns the query commands for IBC connections
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "IBC connection query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdQueryConnections(cdc, reg),
		GetCmdQueryConnection(cdc, reg),
		GetCmdQueryClientConnections(cdc, reg),
	)

	return queryCmd
}
