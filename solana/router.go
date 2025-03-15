package solana

import (
	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/puper/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/puper/tx-parser/solana/types"
)

func router(meta *rpc.TransactionMeta, txn *solanago.Transaction, instructionIdx int) (action types.Action, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			action = nil
		}
	}()
	programID := txn.Message.AccountKeys[txn.Message.Instructions[instructionIdx].ProgramIDIndex].String()
	instruction := txn.Message.Instructions[instructionIdx]
	switch programID {
	/**
	case systemProgram.Program:
		return SystemProgramParsers.InstructionRouter(result, instruction)
	case tokenProgram.Program:
		return TokenProgramParsers.InstructionRouter(result, instruction)
	case computeBudget.Program:
		return ComputeBudgetParsers.InstructionRouter(result, instruction)
	*/
	case pumpfun.Program:
		return PumpfunParsers.InstructionRouter(meta, txn, instruction)
	/*
		case jupiterDCA.Program:
			return JupiterDCA.InstructionRouter(result, instruction)
		case raydiumLiquidityPoolV4.Program:
			return RaydiumLiquidityPoolV4.InstructionRouter(result, instruction, instructionIdx)
		case jupiterAggregatorV6.Program:
			return JupiterAggregatorV6.InstructionRouter(result, instruction)
		case OKXDEXAggregationRouterV2.Program:
			return OKXDEXAggregationRouterV2Parsers.InstructionRouter(result, instruction)
	*/
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       programID,
				ProgramName:     "Unknown",
				InstructionName: "Unknown",
			},
		}, nil
	}
}
