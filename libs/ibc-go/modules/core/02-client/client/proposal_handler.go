package client

import (
	"net/http"

	cliContext "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/client/context"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types/rest"
	"github.com/zhengjianfeng1103/fbc/libs/ibc-go/modules/core/02-client/client/cli"
	govclient "github.com/zhengjianfeng1103/fbc/x/gov/client"
	govrest "github.com/zhengjianfeng1103/fbc/x/gov/client/rest"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal, emptyRestHandler)
)

func emptyRestHandler(ctx cliContext.CLIContext) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-ibc-client",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for IBC proposals")
		},
	}
}
