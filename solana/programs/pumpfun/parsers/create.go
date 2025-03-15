package parsers

import (
	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	"github.com/puper/tx-parser/solana/types"
)

type CreateData struct {
	Discriminator uint64
	Name          string
	Symbol        string
	Uri           string
}

func CreateParser(meta *rpc.TransactionMeta, txn *solanago.Transaction, instruction solanago.CompiledInstruction, decodedData []byte) (*types.PumpFunCreateAction, error) {
	var createData CreateData
	err := borsh.Deserialize(&createData, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.PumpFunCreateAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Create",
		},
		Who:                    txn.Message.AccountKeys[instruction.Accounts[7]].String(),
		Mint:                   txn.Message.AccountKeys[instruction.Accounts[0]].String(),
		MintAuthority:          txn.Message.AccountKeys[instruction.Accounts[1]].String(),
		BondingCurve:           txn.Message.AccountKeys[instruction.Accounts[2]].String(),
		AssociatedBondingCurve: txn.Message.AccountKeys[instruction.Accounts[3]].String(),
		MplTokenMetadata:       txn.Message.AccountKeys[instruction.Accounts[5]].String(),
		MetaData:               txn.Message.AccountKeys[instruction.Accounts[6]].String(),
		Name:                   createData.Name,
		Symbol:                 createData.Symbol,
		Uri:                    createData.Uri,
	}

	return &action, nil
}
