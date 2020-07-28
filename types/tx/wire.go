package tx

import (
	"github.com/Robbin-Liu/go-binance-sdk/types/msg"
	"github.com/tendermint/go-amino"
	//"github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino" //@tododb
)

// cdc global variable
var Cdc = amino.NewCodec()

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	msg.RegisterCodec(cdc)
}

func init() {
	//cryptoAmino.RegisterAmino(Cdc)
	cryptoamino.RegisterAmino(Cdc) //@tododb
	RegisterCodec(Cdc)
}
