package parsers

import (
	"github.com/mr-tron/base58"
	"github.com/puper/tx-parser/solana/programs/computeBudget"
	"github.com/puper/tx-parser/solana/types"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]

	switch discriminator {
	case computeBudget.SetComputeUnitLimitDiscriminator:
		return SetComputeUnitLimitParser(result, instruction, decode)
	case computeBudget.SetComputeUnitPriceDiscriminator:
		return SetComputeUnitPriceParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     computeBudget.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
