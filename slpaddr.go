package bchutil

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/ripemd160"
)

var SlpPrefixes  = map[string]string{
	chaincfg.MainNetParams.Name: "simpleledger",
	chaincfg.TestNet3Params.Name: "slptest",
}

func CheckEncodeSlpAddress(input []byte, prefix string, t AddressType) string {
	k, err := packAddressData(t, input)
	if err != nil {
		fmt.Println("%v", err)
		return ""
	}
	return Encode(prefix, k)
}

// encodeAddress returns a human-readable payment address given a ripemd160 hash
// and prefix which encodes the bitcoin cash network and address type.  It is used
// in both pay-to-pubkey-hash (P2PKH) and pay-to-script-hash (P2SH) address
// encoding.
func encodeSlpAddress(hash160 []byte, prefix string, t AddressType) string {
	return CheckEncodeSlpAddress(hash160[:ripemd160.Size], prefix, t)
}

// AddressPubKeyHash is an Address for a pay-to-pubkey-hash (P2PKH)
// transaction.
type SlpAddressPubKeyHash struct {
	hash   [ripemd160.Size]byte
	prefix string
}

// NewAddressPubKeyHash returns a new AddressPubKeyHash.  pkHash mustbe 20
// bytes.
func NewSlpAddressPubKeyHash(pkHash []byte, net *chaincfg.Params) (*SlpAddressPubKeyHash, error) {
	return newSlpAddressPubKeyHash(pkHash, net)
}

// newAddressPubKeyHash is the internal API to create a pubkey hash address
// with a known leading identifier byte for a network, rather than looking
// it up through its parameters.  This is useful when creating a new address
// structure from a string encoding where the identifer byte is already
// known.
func newSlpAddressPubKeyHash(pkHash []byte, net *chaincfg.Params) (*SlpAddressPubKeyHash, error) {
	// Check for a valid pubkey hash length.
	if len(pkHash) != ripemd160.Size {
		return nil, errors.New("pkHash must be 20 bytes")
	}

	prefix, ok := SlpPrefixes[net.Name]
	if !ok {
		return nil, errors.New("unknown network parameters")
	}

	addr := &SlpAddressPubKeyHash{prefix: prefix}
	copy(addr.hash[:], pkHash)
	return addr, nil
}

// EncodeAddress returns the string encoding of a pay-to-pubkey-hash
// address.  Part of the Address interface.
func (a *SlpAddressPubKeyHash) EncodeAddress() string {
	return encodeSlpAddress(a.hash[:], a.prefix, P2PKH)
}

// ScriptAddress returns the bytes to be included in a txout script to pay
// to a pubkey hash.  Part of the Address interface.
func (a *SlpAddressPubKeyHash) ScriptAddress() []byte {
	return a.hash[:]
}

// IsForNet returns whether or not the pay-to-pubkey-hash address is associated
// with the passed bitcoin cash network.
func (a *SlpAddressPubKeyHash) IsForNet(net *chaincfg.Params) bool {
	checkPre, ok := SlpPrefixes[net.Name]
	if !ok {
		return false
	}
	return a.prefix == checkPre
}

// String returns a human-readable string for the pay-to-pubkey-hash address.
// This is equivalent to calling EncodeAddress, but is provided so the type can
// be used as a fmt.Stringer.
func (a *SlpAddressPubKeyHash) String() string {
	return a.EncodeAddress()
}

// Hash160 returns the underlying array of the pubkey hash.  This can be useful
// when an array is more appropiate than a slice (for example, when used as map
// keys).
func (a *SlpAddressPubKeyHash) Hash160() *[ripemd160.Size]byte {
	return &a.hash
}



// AddressScriptHash is an Address for a pay-to-script-hash (P2SH)
// transaction.
type SlpAddressScriptHash struct {
	hash  [ripemd160.Size]byte
	prefix string
}

// NewAddressScriptHash returns a new AddressScriptHash.
func NewSlpAddressScriptHash(serializedScript []byte, net *chaincfg.Params) (*SlpAddressScriptHash, error) {
	scriptHash := btcutil.Hash160(serializedScript)
	return newSlpAddressScriptHashFromHash(scriptHash, net)
}

// NewAddressScriptHashFromHash returns a new AddressScriptHash.  scriptHash
// must be 20 bytes.
func NewSlpAddressScriptHashFromHash(scriptHash []byte, net *chaincfg.Params) (*SlpAddressScriptHash, error) {
	return newSlpAddressScriptHashFromHash(scriptHash, net)
}

// newAddressScriptHashFromHash is the internal API to create a script hash
// address with a known leading identifier byte for a network, rather than
// looking it up through its parameters.  This is useful when creating a new
// address structure from a string encoding where the identifer byte is already
// known.
func newSlpAddressScriptHashFromHash(scriptHash []byte, net *chaincfg.Params) (*SlpAddressScriptHash, error) {
	// Check for a valid script hash length.
	if len(scriptHash) != ripemd160.Size {
		return nil, errors.New("scriptHash must be 20 bytes")
	}

	pre, ok := SlpPrefixes[net.Name]
	if !ok {
		return nil, errors.New("unknown network parameters")
	}

	addr := &SlpAddressScriptHash{prefix: pre}
	copy(addr.hash[:], scriptHash)
	return addr, nil
}

// EncodeAddress returns the string encoding of a pay-to-script-hash
// address.  Part of the Address interface.
func (a *SlpAddressScriptHash) EncodeAddress() string {
	return encodeSlpAddress(a.hash[:], a.prefix, P2SH)
}

// ScriptAddress returns the bytes to be included in a txout script to pay
// to a script hash.  Part of the Address interface.
func (a *SlpAddressScriptHash) ScriptAddress() []byte {
	return a.hash[:]
}

// IsForNet returns whether or not the pay-to-script-hash address is associated
// with the passed bitcoin cash network.
func (a *SlpAddressScriptHash) IsForNet(net *chaincfg.Params) bool {
	pre, ok := SlpPrefixes[net.Name]
	if !ok {
		return false
	}
	return pre == a.prefix
}

// String returns a human-readable string for the pay-to-script-hash address.
// This is equivalent to calling EncodeAddress, but is provided so the type can
// be used as a fmt.Stringer.
func (a *SlpAddressScriptHash) String() string {
	return a.EncodeAddress()
}

// Hash160 returns the underlying array of the script hash.  This can be useful
// when an array is more appropiate than a slice (for example, when used as map
// keys).
func (a *SlpAddressScriptHash) Hash160() *[ripemd160.Size]byte {
	return &a.hash
}
