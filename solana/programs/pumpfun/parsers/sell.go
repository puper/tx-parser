package parsers

import (
	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/globals"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	"github.com/puper/tx-parser/solana/types"
)

type SellData struct {
	Discriminator uint64
	Amount        uint64
	MinSolOutput  uint64
}

func SellParser(meta *rpc.TransactionMeta, txn *solanago.Transaction, instruction solanago.CompiledInstruction, decodedData []byte) (*types.PumpFunSellAction, error) {
	var sellData SellData
	err := borsh.Deserialize(&sellData, decodedData)
	if err != nil {
		return nil, err
	}

	var instructionIndex uint16
	for idx, instr := range txn.Message.Instructions {
		if txn.Message.AccountKeys[instr.ProgramIDIndex].String() == pumpfun.Program && instr.Data.String() == instruction.Data.String() {
			instructionIndex = uint16(idx)
			break
		}
	}

	var instructions []solanago.CompiledInstruction
	for _, innerInstruction := range meta.InnerInstructions {
		if innerInstruction.Index == instructionIndex {
			instructions = innerInstruction.Instructions
			break
		}
	}

	sellTokenAmount := uint64(0)
	sellSolAmount := uint64(0)

	for _, instr := range instructions {
		programId := txn.Message.AccountKeys[instr.ProgramIDIndex].String()
		if programId == pumpfun.Program {
			data := instr.Data
			decode := data
			if err != nil {
				return nil, err
			}
			discriminator := *(*[16]byte)(decode[:16])
			mergedDiscriminator := make([]byte, 0, 16)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogDiscriminator[:]...)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogSwapDiscriminator[:]...)
			if discriminator == *(*[16]byte)(mergedDiscriminator[:]) {
				action, err := AnchorSelfCPILogSwapParser(decode)
				if err == nil {
					sellTokenAmount = action.TokenAmount
					sellSolAmount = action.SolAmount
				}
			}
		}
	}

	action := types.PumpFunSellAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Sell",
		},
		Who:             txn.Message.AccountKeys[instruction.Accounts[6]].String(),
		FromToken:       txn.Message.AccountKeys[instruction.Accounts[2]].String(),
		ToToken:         globals.WSOL,
		FromTokenAmount: sellTokenAmount,
		ToTokenAmount:   sellSolAmount,
		MinSolOutput:    sellData.MinSolOutput,
	}
	return &action, nil
}
