package parsers

import (
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/programs/computeBudget"
	"github.com/puper/tx-parser/solana/types"
)

type SetComputeUnitLimitData struct {
	Discriminator uint8
	Unit          uint32
}

func SetComputeUnitLimitParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.ComputeBudgetSetComputeUnitLimitAction, error) {
	var data SetComputeUnitLimitData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.ComputeBudgetSetComputeUnitLimitAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     computeBudget.ProgramName,
			InstructionName: "SetComputeUnitLimit",
		},
		ComputeUnitLimit: data.Unit,
	}

	return &action, nil
}
