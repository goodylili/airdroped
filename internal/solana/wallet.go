package solana

import "github.com/blocto/solana-go-sdk/types"

func GenerateWallet() (string, string) {
	account := types.NewAccount()
	privateKey := account.PrivateKey
	publicKey := account.PublicKey.ToBase58()
	return string(privateKey), publicKey
}
