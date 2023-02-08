package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
	govutils "github.com/zhengjianfeng1103/fbc/x/gov/client/utils"
)

func parseSubmitProposalFlags() (*proposal, error) {
	proposal := &proposal{}
	file := viper.GetString(flagProposal)

	if file == "" {
		proposal.Title = viper.GetString(flagTitle)
		proposal.Description = viper.GetString(flagDescription)
		proposal.Type = govutils.NormalizeProposalType(viper.GetString(flagProposalType))
		proposal.Deposit = viper.GetString(flagDeposit)
		return proposal, nil
	}

	for _, flag := range proposalFlags {
		if viper.GetString(flag) != "" {
			return nil, fmt.Errorf("--%s flag provided alongside --proposal, which is a noop", flag)
		}
	}

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, proposal)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}
