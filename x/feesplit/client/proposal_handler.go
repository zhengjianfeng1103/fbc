package client

import (
	"github.com/zhengjianfeng1103/fbc/x/feesplit/client/cli"
	"github.com/zhengjianfeng1103/fbc/x/feesplit/client/rest"
	govcli "github.com/zhengjianfeng1103/fbc/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
