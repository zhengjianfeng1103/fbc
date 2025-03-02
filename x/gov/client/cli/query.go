package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/FiboChain/fbc/libs/cosmos-sdk/client"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/client/context"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/version"
	"github.com/FiboChain/fbc/x/gov/client/utils"

	"github.com/FiboChain/fbc/libs/cosmos-sdk/client/flags"
	"github.com/FiboChain/fbc/x/gov/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group gov queries under a subcommand
	govQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the governance module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	govQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryProposal(queryRoute, cdc),
		GetCmdQueryProposals(queryRoute, cdc),
		getCmdQueryVote(queryRoute, cdc),
		getCmdQueryVotes(queryRoute, cdc),
		GetCmdQueryParam(queryRoute, cdc),
		GetCmdQueryParams(queryRoute, cdc),
		GetCmdQueryProposer(queryRoute, cdc),
		getCmdQueryDeposit(queryRoute, cdc),
		getCmdQueryDeposits(queryRoute, cdc),
		GetCmdQueryTally(queryRoute, cdc))...)

	return govQueryCmd
}

// GetCmdQueryProposal implements the query proposal command.
func GetCmdQueryProposal(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "proposal [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a single proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details for a proposal. You can find the
proposal-id by running "%s query gov proposals".

Example:
$ %s query gov proposal 1
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
			}

			// Query the proposal
			res, err := utils.QueryProposalByID(proposalID, cliCtx, queryRoute)
			if err != nil {
				return err
			}

			var proposal types.Proposal
			cdc.MustUnmarshalJSON(res, &proposal)
			return cliCtx.PrintOutput(proposal) // nolint:errcheck
		},
	}
}

// GetCmdQueryProposals implements a query proposals command.
func GetCmdQueryProposals(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposals",
		Short: "Query proposals with optional filters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for a all proposals. You can filter the returns with the following flags.

Example:
$ %s query gov proposals --depositor ex1rf9wr069pt64e58f2w3mjs9w72g8vemzw26658
$ %s query gov proposals --voter ex1rf9wr069pt64e58f2w3mjs9w72g8vemzw26658
$ %s query gov proposals --status (DepositPeriod|VotingPeriod|Passed|Rejected)
`,
				version.ClientName, version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			bechDepositorAddr := viper.GetString(flagDepositor)
			bechVoterAddr := viper.GetString(flagVoter)
			strProposalStatus := viper.GetString(flagStatus)
			numLimit := uint64(viper.GetInt64(flagNumLimit))

			var depositorAddr sdk.AccAddress
			var voterAddr sdk.AccAddress
			var proposalStatus types.ProposalStatus

			params := types.NewQueryProposalsParams(proposalStatus, numLimit, voterAddr, depositorAddr)

			if len(bechDepositorAddr) != 0 {
				depositorAddr, err := sdk.AccAddressFromBech32(bechDepositorAddr)
				if err != nil {
					return err
				}
				params.Depositor = depositorAddr
			}

			if len(bechVoterAddr) != 0 {
				voterAddr, err := sdk.AccAddressFromBech32(bechVoterAddr)
				if err != nil {
					return err
				}
				params.Voter = voterAddr
			}

			if len(strProposalStatus) != 0 {
				proposalStatus, err := types.ProposalStatusFromString(utils.NormalizeProposalStatus(strProposalStatus))
				if err != nil {
					return err
				}
				params.ProposalStatus = proposalStatus
			}

			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/proposals", queryRoute), bz)
			if err != nil {
				return err
			}

			var matchingProposals types.Proposals
			err = cdc.UnmarshalJSON(res, &matchingProposals)
			if err != nil {
				return err
			}

			if len(matchingProposals) == 0 {
				return fmt.Errorf("No matching proposals found")
			}

			return cliCtx.PrintOutput(matchingProposals) // nolint:errcheck
		},
	}

	cmd.Flags().String(flagNumLimit, "", "(optional) limit to latest [number] proposals. Defaults to all proposals")
	cmd.Flags().String(flagDepositor, "", "(optional) filter by proposals deposited on by depositor")
	cmd.Flags().String(flagVoter, "", "(optional) filter by proposals voted on by voted")
	cmd.Flags().String(flagStatus, "", "(optional) filter proposals by proposal status, status: deposit_period/voting_period/passed/rejected")

	return cmd
}

// GetCmdQueryProposer implements the query proposer command.
func GetCmdQueryProposer(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "proposer [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the proposer of a governance proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query which address proposed a proposal with a given ID.

Example:
$ %s query gov proposer 1
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// validate that the proposalID is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s is not a valid uint", args[0])
			}

			prop, err := utils.QueryProposerByTxQuery(cliCtx, proposalID)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(prop)
		},
	}
}

// GetCmdQueryProposal implements the query proposal command.
func GetCmdQueryParam(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "param [param-type]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the parameters (voting|tallying|deposit) of the governance process",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the all the parameters for the governance process.

Example:
$ %s query gov param voting
$ %s query gov param tallying
$ %s query gov param deposit
`,
				version.ClientName, version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query store
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/params/%s", queryRoute, args[0]), nil)
			if err != nil {
				return err
			}
			var out fmt.Stringer
			switch args[0] {
			case "voting":
				var param types.VotingParams
				cdc.MustUnmarshalJSON(res, &param)
				out = param
			case "tallying":
				var param types.TallyParams
				cdc.MustUnmarshalJSON(res, &param)
				out = param
			case "deposit":
				var param types.DepositParams
				cdc.MustUnmarshalJSON(res, &param)
				out = param
			default:
				return fmt.Errorf("Argument must be one of (voting|tallying|deposit), was %s", args[0])
			}

			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryProposal implements the query proposal command.
func GetCmdQueryParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the parameters of the governance process",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the all the parameters for the governance process.

Example:
$ %s query gov params
`,
				version.ClientName,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			tp, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/params/tallying", queryRoute), nil)
			if err != nil {
				return err
			}
			dp, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/params/deposit", queryRoute), nil)
			if err != nil {
				return err
			}
			vp, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/params/voting", queryRoute), nil)
			if err != nil {
				return err
			}

			var tallyParams types.TallyParams
			cdc.MustUnmarshalJSON(tp, &tallyParams)
			var depositParams types.DepositParams
			cdc.MustUnmarshalJSON(dp, &depositParams)
			var votingParams types.VotingParams
			cdc.MustUnmarshalJSON(vp, &votingParams)

			return cliCtx.PrintOutput(types.NewParams(votingParams, tallyParams, depositParams))
		},
	}
}

// Command to Get a Proposal Information
// getCmdQueryVote implements the query proposal vote command.
func getCmdQueryVote(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote [proposal-id] [voter-addr]",
		Args:  cobra.ExactArgs(2),
		Short: "Query details of a single vote",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details for a single vote on a proposal given its identifier.

Example:
$ %s query gov vote 1 ex1cftp8q8g4aa65nw9s5trwexe77d9t6cr8ndu02
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			voterAddr, proposalID, _, err := parse(cliCtx, queryRoute, args)
			if err != nil {
				return err
			}

			params := types.NewQueryVoteParams(proposalID, voterAddr)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/vote", queryRoute), bz)
			if err != nil {
				return err
			}

			var vote types.Vote
			if err := cdc.UnmarshalJSON(res, &vote); err != nil {
				return err
			}

			if vote.Empty() {
				res, err = utils.QueryVoteByTxQuery(cliCtx, params)
				if err != nil {
					return err
				}
				if err := cdc.UnmarshalJSON(res, &vote); err != nil {
					return err
				}
			}
			return cliCtx.PrintOutput(vote) //nolint:errcheck
		},
	}
}

func getDepositsOrVotes(cdc *codec.Codec, queryRoute string, args []string, isDeposits bool) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	_, proposalID, res, err := parse(cliCtx, queryRoute, args)
	if err != nil {
		return err
	}
	params := types.NewQueryProposalParams(proposalID)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return err
	}

	var proposal types.Proposal
	cdc.MustUnmarshalJSON(res, &proposal)

	propStatus := proposal.Status
	if isDeposits {
		if !(propStatus == types.StatusVotingPeriod || propStatus == types.StatusDepositPeriod) {
			res, err = utils.QueryDepositsByTxQuery(cliCtx, params)
		} else {
			res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/deposits", queryRoute), bz)
		}
	} else {
		if !(propStatus == types.StatusVotingPeriod || propStatus == types.StatusDepositPeriod) {
			res, err = utils.QueryVotesByTxQuery(cliCtx, params)
		} else {
			res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/votes", queryRoute), bz)
		}
	}

	if err != nil {
		return err
	}

	type DepositVotes = types.Votes
	if isDeposits {
		type DepositVotes = types.Deposits
	}
	var dep DepositVotes
	cdc.MustUnmarshalJSON(res, &dep)
	return cliCtx.PrintOutput(dep)
}

// getCmdQueryVotes implements the command to query for proposal votes.
func getCmdQueryVotes(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "votes [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query votes on a proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query vote details for a single proposal by its identifier.

Example:
$ %s query gov votes 1
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDepositsOrVotes(cdc, queryRoute, args, false)
		},
	}
}

// Command to Get a specific deposit Information
// getCmdQueryDeposit implements the query proposal deposit command.
func getCmdQueryDeposit(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [proposal-id] [depositer-addr]",
		Args:  cobra.ExactArgs(2),
		Short: "Query details of a deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details for a single proposal deposit on a proposal by its identifier.

Example:
$ %s query gov deposit 1 ex1rf9wr069pt64e58f2w3mjs9w72g8vemzw26658
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			depositorAddr, proposalID, _, err := parse(cliCtx, queryRoute, args)
			if err != nil {
				return err
			}

			params := types.NewQueryDepositParams(proposalID, depositorAddr)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/deposit", queryRoute), bz)
			if err != nil {
				return err
			}

			var deposit types.Deposit
			cdc.MustUnmarshalJSON(res, &deposit)

			if deposit.Empty() {
				res, err = utils.QueryDepositByTxQuery(cliCtx, params)
				if err != nil {
					return err
				}
				cdc.MustUnmarshalJSON(res, &deposit)
			}

			return cliCtx.PrintOutput(deposit)
		},
	}
}

// getCmdQueryDeposits implements the command to query for proposal deposits.
func getCmdQueryDeposits(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposits [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query deposits on a proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details for all deposits on a proposal.
You can find the proposal-id by running "%s query gov proposals".

Example:
$ %s query gov deposits 1
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getDepositsOrVotes(cdc, queryRoute, args, true)
		},
	}
}

func parse(cliCtx context.CLIContext, queryRoute string, args []string) (sdk.AccAddress, uint64, []byte, error) {
	// validate that the proposal id is a uint
	proposalID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return nil, 0, []byte{},
			fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
	}

	// check to see if the proposal is in the store
	res, err := utils.QueryProposalByID(proposalID, cliCtx, queryRoute)
	if err != nil {
		return nil, proposalID, []byte{}, fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
	}

	var addr sdk.AccAddress
	if len(args) > 1 {
		addr, err = sdk.AccAddressFromBech32(args[1])
		if err != nil {
			return addr, proposalID, res, fmt.Errorf("invalid address：%s", args[1])
		}
	}
	return addr, proposalID, res, nil
}

// GetCmdQueryTally implements the command to query for proposal tally result.
func GetCmdQueryTally(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "tally [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Get the tally of a proposal vote",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query tally of votes on a proposal. You can find
the proposal-id by running "%s query gov proposals".

Example:
$ %s query gov tally 1
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			// check to see if the proposal is in the store
			_, err = utils.QueryProposalByID(proposalID, cliCtx, queryRoute)
			if err != nil {
				return fmt.Errorf("failed to fetch proposal-id %d: %s", proposalID, err)
			}

			// Construct query
			params := types.NewQueryProposalParams(proposalID)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			// Query store
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/tally", queryRoute), bz)
			if err != nil {
				return err
			}

			var tally types.TallyResult
			cdc.MustUnmarshalJSON(res, &tally)
			return cliCtx.PrintOutput(tally)
		},
	}
}

// DONTCOVER
