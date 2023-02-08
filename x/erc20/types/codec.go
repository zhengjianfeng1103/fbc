package types

import (
	"github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/codec"
)

// ModuleCdc defines the erc20 module's codec
var ModuleCdc = codec.New()

const (
	TokenMappingProposalName          = "fbexchain/erc20/TokenMappingProposal"
	ProxyContractRedirectProposalName = "fbexchain/erc20/ProxyContractRedirectProposal"
	ContractTemplateProposalName      = "fbexchain/erc20/ContractTemplateProposal"
	CompiledContractProposalName      = "fbexchain/erc20/Contract"
)

// RegisterCodec registers all the necessary types and interfaces for the
// erc20 module
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(TokenMappingProposal{}, TokenMappingProposalName, nil)

	cdc.RegisterConcrete(ProxyContractRedirectProposal{}, ProxyContractRedirectProposalName, nil)
	cdc.RegisterConcrete(ContractTemplateProposal{}, ContractTemplateProposalName, nil)
	cdc.RegisterConcrete(CompiledContract{}, CompiledContractProposalName, nil)
}

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
