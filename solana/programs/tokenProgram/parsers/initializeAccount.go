package parsers

import (
	"github.com/puper/tx-parser/solana/programs/tokenProgram"
	"github.com/puper/tx-parser/solana/types"
)

func InitializeAccountParser(result *types.ParsedResult, instruction types.Instruction) (*types.TokenProgramInitializeAccountAction, error) {

	action := types.TokenProgramInitializeAccountAction{
		BaseAction: types.BaseAction{
			ProgramID:       tokenProgram.Program,
			ProgramName:     tokenProgram.ProgramName,
			InstructionName: "InitializeAccount",
		},
		Account:    result.AccountList[instruction.Accounts[0]],
		Mint:       result.AccountList[instruction.Accounts[1]],
		Owner:      result.AccountList[instruction.Accounts[2]],
		RentSysvar: result.AccountList[instruction.Accounts[3]],
	}
	return &action, nil
}
