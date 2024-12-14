package main

import (
	configurations "airdroped/cmd"
	"airdroped/internal/evm"
	"airdroped/internal/services"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func scaleFloatsToBigInts(values []float64, power uint8) []*big.Int {
	result := make([]*big.Int, len(values))

	// Calculate 10^power manually
	scaleFactor := new(big.Float).SetPrec(256).SetFloat64(1)
	for i := uint8(0); i < power; i++ {
		scaleFactor = scaleFactor.Mul(scaleFactor, big.NewFloat(10))
	}

	for i, val := range values {
		bigVal := new(big.Float).SetFloat64(val)
		bigVal = bigVal.Mul(bigVal, scaleFactor)
		intVal, _ := bigVal.Int(nil) // Convert to big.Int, discarding fractional part
		result[i] = intVal
	}
	return result
}

func main() {
	config := configurations.LoadConfigurations()
	addresses, amounts, err := services.ReadColumnByNameAndValidate(config.CSVFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize Ethereum client
	ethClient, err := evm.NewEthClient(config.RPCURL, config.PrivateKey, config.ChainID, config.CA)
	if err != nil {
		log.Fatalf("Error initializing Ethereum client: %v", err)
	}
	// Get token decimals
	decimals, err := ethClient.GetTokenDecimals(config.TokenAddress)
	if err != nil {
		log.Fatalf("Error getting token decimals: %v", err)
	}
	// Check balance
	balance, err := ethClient.CheckBalance(config.TokenAddress, config.PublicKey)
	if err != nil {
		log.Fatalf("Error checking balance: %v", err)
	}
	fmt.Printf("Current balance: %s\n", balance.String())

	scaledAmounts := scaleFloatsToBigInts(amounts, decimals)

	err = ethClient.ApproveToken(
		config.PrivateKey,
		common.HexToAddress(config.TokenAddress),
		common.HexToAddress(config.CA),
		balance,
	)
	if err != nil {
		log.Fatalf("Error approving token transfer: %v", err)
	}
	fmt.Println("Token transfer approved successfully")

	// Then perform the batch transfer
	fmt.Println("Initiating batch transfer...")
	if err := ethClient.BatchTransferAssets(config.TokenAddress, addresses, scaledAmounts, config.ChainID); err != nil {
		log.Fatalf("Batch transfer failed: %v", err)
	}
	fmt.Println("Batch transfer completed successfully")
}
