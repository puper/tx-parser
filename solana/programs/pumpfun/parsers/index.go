package parsers

import (
	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	"github.com/puper/tx-parser/solana/types"
)

func InstructionRouter(meta *rpc.TransactionMeta, txn *solanago.Transaction, instruction solanago.CompiledInstruction) (types.Action, error) {
	data := instruction.Data
	decode := data
	discriminator := *(*[8]byte)(decode[:8])

	switch discriminator {
	case pumpfun.BuyDiscriminator:
		return BuyParser(meta, txn, instruction, decode)
	case pumpfun.SellDiscriminator:
		return SellParser(meta, txn, instruction, decode)
	case pumpfun.CreateDiscriminator:
		return CreateParser(meta, txn, instruction, decode)
	case pumpfun.AnchorSelfCPILogDiscriminator:
		subDiscriminator := *(*[8]byte)(decode[8:16])
		switch subDiscriminator {
		case pumpfun.AnchorSelfCPILogSwapDiscriminator:
			return AnchorSelfCPILogSwapParser(decode)
		default:
			return types.UnknownAction{
				BaseAction: types.BaseAction{
					ProgramID:       txn.Message.AccountKeys[instruction.ProgramIDIndex].String(),
					ProgramName:     pumpfun.ProgramName,
					InstructionName: "AnchorSelfCPILog Unknown",
				},
			}, nil
		}

	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       txn.Message.AccountKeys[instruction.ProgramIDIndex].String(),
				ProgramName:     pumpfun.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
