package main

import (
	"airdroped/internal/evm"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const BatchContractAddress = "0xbAE75Da35Ea0CA0C19c7A78B5807c2b6eb3a9427"

// EthClient is a struct encapsulating Ethereum client operations
type EthClient struct {
	Client       *ethclient.Client
	Auth         *bind.TransactOpts
	ChainID      *big.Int
	PrivateKey   string
	ContractAddr common.Address
}

// NewEthClient initializes a new EthClient
func NewEthClient(rpcURL, privateKey string, chainID *big.Int, contractAddr string) (*EthClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %w", err)
	}

	return &EthClient{
		Client:       client,
		Auth:         auth,
		ChainID:      chainID,
		PrivateKey:   privateKey,
		ContractAddr: common.HexToAddress(contractAddr),
	}, nil
}

// BatchTransfer executes a batch transfer
func (e *EthClient) BatchTransfer(tokenAddress string, recipients []string, amounts []*big.Int) error {
	contract, err := evm.NewBatchtransfer(e.ContractAddr, e.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate contract: %w", err)
	}

	token := common.HexToAddress(tokenAddress)
	var addresses []common.Address
	for _, recipient := range recipients {
		addresses = append(addresses, common.HexToAddress(recipient))
	}

	var scaledAmounts []*big.Int
	for _, amount := range amounts {
		scaledAmounts = append(scaledAmounts, amount)
	}

	tx, err := contract.BatchTransfer(e.Auth, token, addresses, scaledAmounts)
	if err != nil {
		return fmt.Errorf("failed to execute batch transfer: %w", err)
	}

	fmt.Printf("Batch transfer transaction sent: %s\n", tx.Hash().Hex())
	return nil
}

func main() {
	// Configuration
	rpcURL := "https://arbitrum-mainnet.infura.io/v3/23c96151438348e79d1e963324d33b02"
	privateKey := "f0029d055d03c9e06e4c87ac6f0c6209948707987765d47e6121e873d9fae3cf"
	chainID := big.NewInt(42161)

	// Initialize EthClient
	client, err := NewEthClient(rpcURL, privateKey, chainID, BatchContractAddress)
	if err != nil {
		log.Fatalf("Error initializing Ethereum client: %v", err)
	}

	// Example batch transfer
	tokenAddress := "0xaf88d065e77c8cc2239327c5edb3a432268e5831"
	recipients := []string{
		"0xfa53d837b5ddd616007f91487f041d27edb683a0",
		"0x34edecd3108391dd044b617e2ec9c150a78aec17",
		"0x047f17fedafb60d4290997a3f17544f3026d6ef3",
	}
	scaleFactor := int64(1)

	// Represent 0.01 using big.Rat
	// Represent 0.01 as an integer (scaled to cents)
	amounts := []*big.Int{
		big.NewInt(int64(0.01 * float64(scaleFactor))), // Scaled amount for recipient 1
		big.NewInt(int64(0.01 * float64(scaleFactor))), // Scaled amount for recipient 2
		big.NewInt(int64(0.01 * float64(scaleFactor))), // Scaled amount for recipient 3
	}
	if err := client.BatchTransfer(tokenAddress, recipients, amounts); err != nil {
		log.Fatalf("Batch transfer failed: %v", err)
	}
}
