package fss

import (
	"github.com/spf13/cobra"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/server"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/iavl"
	tmtypes "github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

const (
	flagDataDir = "data_dir"
)

func Command(ctx *server.Context) *cobra.Command {
	iavl.SetLogger(ctx.Logger.With("module", "iavl"))
	return fssCmd
}

var fssCmd = &cobra.Command{
	Use:   "fss",
	Short: "FSS is an auxiliary fast storage system to IAVL",
	Long: `IAVL fast storage related commands:
This command include a set of command of the IAVL fast storage.
include create sub command`,
}

func init() {
	fssCmd.PersistentFlags().StringP(flagDataDir, "d", "./", "The chain data file location")
	fssCmd.PersistentFlags().String(sdk.FlagDBBackend, tmtypes.DBBackend, "Database backend: goleveldb | rocksdb")
}
