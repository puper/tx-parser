package parsers

import (
	"github.com/mr-tron/base58"
	"github.com/puper/tx-parser/solana/programs/raydiumLiquidityPoolV4"
	"github.com/puper/tx-parser/solana/types"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction, instructionIdx int) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]
	switch discriminator {
	case raydiumLiquidityPoolV4.SwapDiscriminator:
		return SwapParser(result, instruction, instructionIdx, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     raydiumLiquidityPoolV4.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
