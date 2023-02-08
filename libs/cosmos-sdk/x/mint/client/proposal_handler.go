package client

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/zhengjianfeng1103/fbc/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)
)
