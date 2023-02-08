package client

import (
	"github.com/zhengjianfeng1103/fbc/x/farm/client/cli"
	"github.com/zhengjianfeng1103/fbc/x/farm/client/rest"
	govcli "github.com/zhengjianfeng1103/fbc/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
