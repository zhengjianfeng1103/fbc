package client

import (
	govclient "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/gov/client"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
