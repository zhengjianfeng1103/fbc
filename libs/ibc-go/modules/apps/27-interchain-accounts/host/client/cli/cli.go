package cli

import (
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	interfacetypes "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
)

// GetQueryCmd returns the query commands for the ICA host submodule
func GetQueryCmd(cdc *codec.CodecProxy, reg interfacetypes.InterfaceRegistry) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "host",
		Short:                      "interchain-accounts host subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		GetCmdParams(cdc, reg),
		GetCmdPacketEvents(cdc, reg),
	)

	return queryCmd
}
