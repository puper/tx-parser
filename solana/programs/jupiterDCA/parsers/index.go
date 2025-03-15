package parsers

import (
	"github.com/mr-tron/base58"
	"github.com/puper/tx-parser/solana/programs/jupiterDCA"
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
	case jupiterDCA.OpenDcaV2Discriminator:
		return OpenDcaV2Parser(result, instruction, decode)
	case jupiterDCA.EndAndCloseDiscriminator:
		return EndAndCloseParser(result, instruction, decode)
	case jupiterDCA.CloseDcaDiscriminator:
		return CloseDcaParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     jupiterDCA.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
