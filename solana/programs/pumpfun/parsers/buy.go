package parsers

import (
	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/globals"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	"github.com/puper/tx-parser/solana/types"
)

type BuyData struct {
	Discriminator uint64
	Amount        uint64
	MaxSolCost    uint64
}

func BuyParser(meta *rpc.TransactionMeta, txn *solanago.Transaction, instruction solanago.CompiledInstruction, decodedData []byte) (*types.PumpFunBuyAction, error) {
	var buyData BuyData
	err := borsh.Deserialize(&buyData, decodedData)
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

	buyTokenAmount := uint64(0)
	buySolAmount := uint64(0)

	for _, instr := range instructions {
		programId := txn.Message.AccountKeys[instr.ProgramIDIndex].String()
		if programId == pumpfun.Program {
			decode := instr.Data
			discriminator := *(*[16]byte)(decode[:16])
			mergedDiscriminator := make([]byte, 0, 16)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogDiscriminator[:]...)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogSwapDiscriminator[:]...)
			if discriminator == *(*[16]byte)(mergedDiscriminator[:]) {
				action, err := AnchorSelfCPILogSwapParser(decode)
				if err == nil {
					buyTokenAmount = action.TokenAmount
					buySolAmount = action.SolAmount
				}
			}
		}
	}

	action := types.PumpFunBuyAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Buy",
		},
		Who:             txn.Message.AccountKeys[instruction.Accounts[6]].String(),
		ToToken:         txn.Message.AccountKeys[instruction.Accounts[2]].String(),
		FromToken:       globals.WSOL,
		ToTokenAmount:   buyTokenAmount,
		FromTokenAmount: buySolAmount,
		MaxSolCost:      buyData.MaxSolCost,
	}

	return &action, nil
}
