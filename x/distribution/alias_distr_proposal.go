// nolint
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/zhengjianfeng1103/fbc/x/distribution/types
// ALIASGEN: github.com/zhengjianfeng1103/fbc/x/distribution/client
package distribution

import (
	"github.com/zhengjianfeng1103/fbc/x/distribution/client"
	"github.com/zhengjianfeng1103/fbc/x/distribution/types"
)

var (
	NewMsgWithdrawDelegatorReward          = types.NewMsgWithdrawDelegatorReward
	CommunityPoolSpendProposalHandler      = client.CommunityPoolSpendProposalHandler
	ChangeDistributionTypeProposalHandler  = client.ChangeDistributionTypeProposalHandler
	WithdrawRewardEnabledProposalHandler   = client.WithdrawRewardEnabledProposalHandler
	RewardTruncatePrecisionProposalHandler = client.RewardTruncatePrecisionProposalHandler
	NewMsgWithdrawDelegatorAllRewards      = types.NewMsgWithdrawDelegatorAllRewards
)
