package parsers

import (
	"github.com/puper/tx-parser/solana/programs/jupiterDCA"
	"github.com/puper/tx-parser/solana/types"
)

func EndAndCloseParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.JupiterDcaEndAndCloseAction, error) {
	return &types.JupiterDcaEndAndCloseAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiterDCA.ProgramName,
			InstructionName: "EndAndClose",
		},
		Keeper:     result.AccountList[instruction.Accounts[0]],
		Dca:        result.AccountList[instruction.Accounts[1]],
		InputMint:  result.AccountList[instruction.Accounts[2]],
		OutputMint: result.AccountList[instruction.Accounts[3]],
		InAta:      result.AccountList[instruction.Accounts[4]],
		OutAta:     result.AccountList[instruction.Accounts[5]],
		User:       result.AccountList[instruction.Accounts[6]],
		UserOutAta: result.AccountList[instruction.Accounts[7]],
	}, nil
}
