package bchutil

import (
	"testing"
	"github.com/btcsuite/btcd/chaincfg"
)

var valid []string = []string {
	"prefix:x64nx6hz",
	"PREFIX:X64NX6HZ",
	"p:gpf8m4h7",
	"bitcoincash:qpzry9x8gf2tvdw0s3jn54khce6mua7lcw20ayyn",
	"bchtest:testnetaddress4d6njnut",
	"bchreg:555555555555555555555555555555555555555555555udxmlmrz",
}

func TestValid(t *testing.T) {
	for _, s := range valid {
		_, _, err := DecodeCashAddr(s)
		if err != nil {
			t.Error(err)
		}
	}
}

var Invalid []string = []string {
	"prefix:x32nx6hz",
	"prEfix:x64nx6hz",
	"prefix:x64nx6Hz",
	"pref1x:6m8cxv73",
	"prefix:",
	":u9wsx07j",
	"bchreg:555555555555555555x55555555555555555555555555udxmlmrz",
	"bchreg:555555555555555555555555555555551555555555555udxmlmrz",
	"pre:fix:x32nx6hz",
	"prefixx64nx6hz",
}

func TestInvalid(t *testing.T) {
	for _, s := range Invalid {
		_, _, err := DecodeCashAddr(s)
		if err == nil {
			t.Error("Failed to error on invalid string")
		}
	}
}

func TestDecodeCashAddress(t *testing.T) {
	// Mainnet
	addr, err := DecodeCashAddress("bitcoincash:qpzry9x8gf2tvdw0s3jn5grmq650p", &chaincfg.MainNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bitcoincash:qpzry9x8gf2tvdw0s3jn5grmq650p" {
		t.Error("Address decoding error")
	}
	// Testnet
	addr, err = DecodeCashAddress("bchtest:qpzry9x8gf2tvdw0s3jn5839nranq", &chaincfg.TestNet3Params)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchtest:qpzry9x8gf2tvdw0s3jn5839nranq" {
		t.Error("Address decoding error")
	}
	// Regtest
	addr, err = DecodeCashAddress("bchreg:qpzry9x8gf2tvdw0s3jn5tgutzngm", &chaincfg.RegressionNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchreg:qpzry9x8gf2tvdw0s3jn5tgutzngm" {
		t.Error("Address decoding error")
	}
}

var dataElement = []byte{203, 72, 18, 50, 41, 156, 213, 116, 49, 81, 172, 75, 45, 99, 174, 25, 142, 123, 176, 169}
// Second address of https://github.com/Bitcoin-UAHF/spec/blob/master/cashaddr.md#examples-of-address-translation
func TestCashAddressPubKeyHash_EncodeAddress(t *testing.T) {
	// Mainnet
	addr, err := NewCashAddressPubKeyHash(dataElement, &chaincfg.MainNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bitcoincash:qr95sy3j9xwd2ap32xkykttr4cvcu7as4y0qverfuy" {
		t.Error("Address decoding error")
	}
	// Testnet
	addr, err = NewCashAddressPubKeyHash(dataElement, &chaincfg.TestNet3Params)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchtest:qpzry9x8gf2tvdw0s3jn5wzx48p70" {
		t.Error("Address decoding error")
	}
	// Regtest
	addr, err = NewCashAddressPubKeyHash(dataElement, &chaincfg.RegressionNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchreg:qpzry9x8gf2tvdw0s3jn5yk7zka9p" {
		t.Error("Address decoding error")
	}
}

var dataElement2 = []byte{118, 160, 64, 83, 189, 160, 168, 139, 218, 81, 119, 184, 106, 21, 195, 178, 159, 85, 152, 115}
// 4th address of https://github.com/Bitcoin-UAHF/spec/blob/master/cashaddr.md#examples-of-address-translation
func TestCashAddressScriptHash_EncodeAddress(t *testing.T) {
	// Mainnet
	addr, err := NewCashAddressScriptHashFromHash(dataElement2, &chaincfg.MainNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq" {
		t.Error("Address decoding error")
	}
	// Testnet
	addr, err = NewCashAddressScriptHashFromHash(dataElement2, &chaincfg.TestNet3Params)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchtest:gpzry9x8gf2tvdw0s3jn5l6cq5ykc" {
		t.Error("Address decoding error")
	}
	// Regtest
	addr, err = NewCashAddressScriptHashFromHash(dataElement2, &chaincfg.RegressionNetParams)
	if err != nil {
		t.Error(err)
	}
	if addr.String() != "bchreg:gpzry9x8gf2tvdw0s3jn54wqh9cdk" {
		t.Error("Address decoding error")
	}
}
