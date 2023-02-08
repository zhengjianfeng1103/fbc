package client

import (
	"github.com/zhengjianfeng1103/fbc/x/dex/client/cli"
	"github.com/zhengjianfeng1103/fbc/x/dex/client/rest"
	govclient "github.com/zhengjianfeng1103/fbc/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
