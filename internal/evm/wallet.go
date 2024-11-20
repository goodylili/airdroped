package evm

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// CreateWallet generates a new Ethereum wallet for user.
func CreateWallet() (address, privateKey string, err error) {
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	privateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address = string(crypto.PubkeyToAddress(*publicKeyECDSA).Hex())

	return address, privateKey, nil
}
