package parsers

import (
	"github.com/mr-tron/base58"
	"github.com/puper/tx-parser/solana/programs/OKXDEXAggregationRouterV2"
	"github.com/puper/tx-parser/solana/types"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := *(*[8]byte)(decode[:8])

	switch discriminator {
	case OKXDEXAggregationRouterV2.CommissionSplProxySwapDiscriminator:
		return CommissionSplProxySwapParser(result, instruction)
	case OKXDEXAggregationRouterV2.SwapDiscriminator:
		return SwapParser(result, instruction, decode)
	case OKXDEXAggregationRouterV2.CommissionSolSwap2Discriminator:
		return CommissionSolSwap2Parser(result, instruction)

	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     OKXDEXAggregationRouterV2.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
