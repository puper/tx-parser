package parsers

import (
	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
	"github.com/puper/tx-parser/solana/programs/pumpfun"
	"github.com/puper/tx-parser/solana/types"
)

type AnchorSelfCPILogSwapData struct {
	Discriminator        [16]byte
	Mint                 [32]byte
	SolAmount            uint64
	TokenAmount          uint64
	IsBuy                bool
	User                 [32]byte
	Timestamp            int64
	VirtualSolReserves   uint64
	VirtualTokenReserves uint64
	IDontKnow1           uint64
	IDontKnow2           uint64
}

func AnchorSelfCPILogSwapParser(decodedData []byte) (*types.PumpFunAnchorSelfCPILogSwapAction, error) {
	var data AnchorSelfCPILogSwapData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.PumpFunAnchorSelfCPILogSwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     "pumpfun",
			InstructionName: "AnchorSelfCPILog Swap",
		},
		Mint:                 base58.Encode(data.Mint[:]),
		SolAmount:            data.SolAmount,
		TokenAmount:          data.TokenAmount,
		IsBuy:                data.IsBuy,
		User:                 base58.Encode(data.User[:]),
		Timestamp:            data.Timestamp,
		VirtualTokenReserves: data.VirtualTokenReserves,
		VirtualSolReserves:   data.VirtualSolReserves,
		IDontKnow1:           data.IDontKnow1,
		IDontKnow2:           data.IDontKnow2,
	}

	return &action, nil
}
