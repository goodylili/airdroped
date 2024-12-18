package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

// ABIs
const (
	erc20ABI = `[
		{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"stateMutability":"view","type":"function"},
		{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},
		{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"}
	]`
)

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

// GetTokenDecimals retrieves the number of decimals for a token
func (e *EthClient) GetTokenDecimals(tokenAddress string) (uint8, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return 0, err
	}

	contractAddr := common.HexToAddress(tokenAddress)
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	result, err := e.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return 0, err
	}

	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", result)
	if err != nil {
		return 0, err
	}

	return decimals, nil
}

// CheckBalance fetches the balance of a user's address for a specific ERC20 token
func (e *EthClient) CheckBalance(tokenAddress, userAddress string) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return new(big.Int), err
	}
	contractAddr := common.HexToAddress(tokenAddress)
	user := common.HexToAddress(userAddress)

	data, err := parsedABI.Pack("balanceOf", user)
	if err != nil {
		return new(big.Int), err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	result, err := e.Client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return new(big.Int), err
	}

	balance := new(big.Int)
	balance.SetBytes(result)

	return balance, nil
}

func (e *EthClient) BatchTransferAssets(tokenAddress string, recipients []string, amounts []*big.Int, chainID *big.Int) error {
	// Create a new authorization with the correct gas settings
	privateKey, err := crypto.HexToECDSA(e.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	// Get the latest nonce
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := e.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	// Create new transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // Set an appropriate gas limit
	auth.GasPrice = gasPrice

	// Create contract instance
	batchContract, err := NewBatchtransfer(e.ContractAddr, e.Client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %w", err)
	}

	recipientAddresses := make([]common.Address, len(recipients))
	for i, addr := range recipients {
		recipientAddresses[i] = common.HexToAddress(addr)
	}

	// Execute the batch transfer with the new auth
	tx, err := batchContract.BatchTransfer(auth, common.HexToAddress(tokenAddress), recipientAddresses, amounts)
	if err != nil {
		return fmt.Errorf("batch transfer failed: %w", err)
	}

	fmt.Printf("Batch transfer initiated. Transaction hash: %s\n", tx.Hash().Hex())
	return nil
}

// ApproveToken approves the spender to transfer tokens
func (e *EthClient) ApproveToken(userWalletPrivateKey string, tokenAddress, spender common.Address, amount *big.Int) error {
	privateKey, err := crypto.HexToECDSA(userWalletPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return err
	}

	data, err := parsedABI.Pack("approve", spender, amount)
	if err != nil {
		return err
	}

	nonce, err := e.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	}

	gasLimit, err := e.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	chainID, err := e.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = e.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	fmt.Printf("Approval transaction sent: %s\n", signedTx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), e.Client, signedTx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("approval transaction failed")
	}

	return nil
}
