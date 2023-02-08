package utils

import (
	"io/ioutil"

	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/x/mint/internal/types"
)

// ManageTreasuresProposalJSON defines a ManageTreasureProposal with a deposit used to parse
// manage treasures proposals from a JSON file.
type ManageTreasuresProposalJSON struct {
	Title       string           `json:"title" yaml:"title"`
	Description string           `json:"description" yaml:"description"`
	Treasures   []types.Treasure `json:"treasures" yaml:"treasures"`
	IsAdded     bool             `json:"is_added" yaml:"is_added"`
	Deposit     sdk.SysCoins     `json:"deposit" yaml:"deposit"`
}

// ParseManageTreasuresProposalJSON parses json from proposal file to ManageTreasuresProposalJSON struct
func ParseManageTreasuresProposalJSON(cdc *codec.Codec, proposalFilePath string) (
	proposal ManageTreasuresProposalJSON, err error) {
	contents, err := ioutil.ReadFile(proposalFilePath)
	if err != nil {
		return
	}

	cdc.MustUnmarshalJSON(contents, &proposal)
	return
}
