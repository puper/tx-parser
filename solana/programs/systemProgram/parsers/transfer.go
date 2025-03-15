package parsers

import (
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/programs/systemProgram"
	"github.com/puper/tx-parser/solana/types"
)

type TransferData struct {
	Discriminator uint32
	Lamports      uint64
}

func TransferParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.SystemProgramTransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.SystemProgramTransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     systemProgram.ProgramName,
			InstructionName: "Transfer",
		},
		From:     result.AccountList[instruction.Accounts[0]],
		To:       result.AccountList[instruction.Accounts[1]],
		Lamports: data.Lamports,
	}

	return &action, nil
}
