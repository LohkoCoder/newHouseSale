package workflow

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
	"hash"
)

type Bytes32 [32]byte

func NewBlake2b() hash.Hash {
	hash, _ := blake2b.New256(nil)
	return hash
}

// Blake2b computes blake2b-256 checksum for given data.
func Blake2b(data ...[]byte) (b32 Bytes32) {
	hash := NewBlake2b()
	for _, b := range data {
		hash.Write(b)
	}
	hash.Sum(b32[:0])
	return
}

func MakeCustomerKey(id, name, phoneNum string) string{

	hashBytes := Blake2b([]byte(id), []byte(name), []byte(phoneNum))
	return hex.EncodeToString(hashBytes[:])
}

func MakeOrderKey(orderId,customerId, name, phoneNum string) string{

	hashBytes := Blake2b([]byte(orderId), []byte(customerId), []byte(name), []byte(phoneNum))
	return hex.EncodeToString(hashBytes[:])
}

/**
represents the customer status during the whole purchase phase
*/
type CustomerStatus string

const (
	Registered CustomerStatus = "Registered"
	ConfirmRegistered CustomerStatus = "Registration Confirmed"
	DownPayment CustomerStatus = "Downpayment Paid"
	ConfirmDownPayment CustomerStatus = "Downpayment Confirmed"
	FullPayment CustomerStatus = "Fullpayment Paid"
	ConfirmFullPayment CustomerStatus = "Fullpayment Confirmed"
)