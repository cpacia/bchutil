package bchutil

import (
	"errors"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
)

func PayToAddrScript(addr btcutil.Address) ([]byte, error) {
	var script []byte
	var err error
	script, err = txscript.PayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	script, err = cashPayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	script, err = bitpayPayToAddrScript(addr)
	if err == nil {
		return script, nil
	}
	return script, errors.New("Unrecognized address format")
}
