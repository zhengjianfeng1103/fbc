package rpc

import (
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/flags"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/lcd"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec/types"
)

// ServeCmd creates a CLI command to start Cosmos REST server with web3 RPC API and
// Cosmos rest-server endpoints
func ServeCmd(cdc *codec.CodecProxy, reg types.InterfaceRegistry) *cobra.Command {
	cmd := lcd.ServeCommand(cdc, reg, RegisterRoutes)
	cmd.Flags().String(flagUnlockKey, "", "Select a key to unlock on the RPC server")
	cmd.Flags().String(FlagWebsocket, "8546", "websocket port to listen to")
	cmd.Flags().StringP(flags.FlagBroadcastMode, "b", flags.BroadcastSync, "Transaction broadcasting mode (sync|async|block)")
	return cmd
}
