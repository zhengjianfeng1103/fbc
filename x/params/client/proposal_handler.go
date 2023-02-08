package client

import (
	govclient "github.com/zhengjianfeng1103/fbc/x/gov/client"
	"github.com/zhengjianfeng1103/fbc/x/params/client/cli"
	"github.com/zhengjianfeng1103/fbc/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler in cmsdk
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
