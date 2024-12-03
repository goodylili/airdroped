package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"strings"
	"time"
)

func (e *AirdropError) Error() string {
	log.Printf("[ERROR] Code: %s | Message: %s | Details: %v", e.Code, e.Message, e.Err)
	return fmt.Sprintf("%s: %s (underlying error: %v)", e.Code, e.Message, e.Err)
}

// NewAirdrop creates a new Airdrop instance
func NewAirdrop(rpcURL, privateKeyHex, tokenAddressHex string) (*Airdrop, error) {
	log.Printf("[INIT] Creating Airdrop instance with RPC URL: %s, Token Address: %s", rpcURL, tokenAddressHex)

	if rpcURL == "" || privateKeyHex == "" || tokenAddressHex == "" {
		log.Printf("[ERROR] Missing required parameters: RPC URL, private key, or token address.")
		return nil, &AirdropError{
			Code:    "INVALID_PARAMS",
			Message: "RPC URL, private key, and token address cannot be empty",
		}
	}

	log.Println("[CONNECT] Establishing connection to Ethereum client...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Printf("[ERROR] Failed to connect to Ethereum client at %s: %v", rpcURL, err)
		return nil, &AirdropError{
			Code:    "CLIENT_CONNECTION_FAILED",
			Message: "Failed to connect to Ethereum client",
			Err:     err,
		}
	}

	log.Println("[SECURITY] Parsing private key...")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("[ERROR] Invalid private key provided: %v", err)
		return nil, &AirdropError{
			Code:    "INVALID_PRIVATE_KEY",
			Message: "Failed to parse private key",
			Err:     err,
		}
	}

	log.Println("[CHAIN] Fetching chain ID...")
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Printf("[ERROR] Unable to fetch chain ID: %v", err)
		return nil, &AirdropError{
			Code:    "CHAIN_ID_FETCH_FAILED",
			Message: "Failed to fetch chain ID",
			Err:     err,
		}
	}

	tokenAddress := common.HexToAddress(tokenAddressHex)
	log.Printf("[TOKEN] Initializing token contract at address: %s", tokenAddressHex)
	token, err := NewERC20(tokenAddress, client)
	if err != nil {
		log.Printf("[ERROR] Failed to initialize token contract at %s: %v", tokenAddressHex, err)
		return nil, &AirdropError{
			Code:    "TOKEN_CONTRACT_INIT_FAILED",
			Message: "Failed to create token contract instance",
			Err:     err,
		}
	}

	log.Println("[SUCCESS] Airdrop instance created successfully.")
	return &Airdrop{
		Client:       client,
		PrivateKey:   privateKey,
		TokenAddress: tokenAddress,
		ChainID:      chainID,
		Token:        token,
	}, nil
}

// ConvertToTokenUnits converts human-readable token amount to precise big.Int units
func ConvertToTokenUnits(token ERC20s, amount float64) (*big.Int, error) {
	// Fetch token decimals
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token decimals: %v", err)
	}

	// Convert amount to big.Int with correct decimal places
	amountFloat := new(big.Float).SetFloat64(amount)
	decimalMultiplier := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))

	// Multiply by decimal multiplier
	scaledAmount := new(big.Float).Mul(amountFloat, decimalMultiplier)

	// Convert to big.Int
	tokenUnits, _ := scaledAmount.Int(nil)

	return tokenUnits, nil
}

// MustConvertToTokenUnits is a wrapper that panics if conversion fails
func MustConvertToTokenUnits(token ERC20s, amount float64) *big.Int {
	tokenUnits, err := ConvertToTokenUnits(token, amount)
	if err != nil {
		panic(err)
	}
	return tokenUnits
}

// ValidateAirdropAllocation checks total airdrop amount against sender's balance
func (a *Airdrop) ValidateAirdropAllocation(allocations map[common.Address]*big.Int) error {
	log.Println("[VALIDATION] Validating airdrop allocations...")
	fromAddress := crypto.PubkeyToAddress(a.PrivateKey.PublicKey)
	log.Printf("[SENDER] Address: %s", fromAddress.Hex())

	totalAirdropAmount := big.NewInt(0)
	for recipient, amount := range allocations {
		log.Printf("[ALLOCATION] Recipient: %s | Amount: %s", recipient.Hex(), amount.String())
		totalAirdropAmount.Add(totalAirdropAmount, amount)
	}
	log.Printf("[TOTAL] Total airdrop amount: %s", totalAirdropAmount.String())

	log.Println("[BALANCE] Fetching sender's token balance...")
	balance, err := a.Token.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch balance for sender %s: %v", fromAddress.Hex(), err)
		return &AirdropError{
			Code:    "BALANCE_CHECK_FAILED",
			Message: "Failed to check sender's token balance",
			Err:     err,
		}
	}
	log.Printf("[BALANCE] Sender's balance: %s", balance.String())

	if balance.Cmp(totalAirdropAmount) < 0 {
		log.Printf("[ERROR] Insufficient balance for airdrop. Required: %s, Available: %s", totalAirdropAmount.String(), balance.String())
		return &AirdropError{
			Code:    "INSUFFICIENT_BALANCE",
			Message: fmt.Sprintf("Required: %s, Available: %s", totalAirdropAmount.String(), balance.String()),
		}
	}

	for recipient, amount := range allocations {
		if amount.Cmp(big.NewInt(0)) <= 0 {
			log.Printf("[ERROR] Invalid allocation amount for recipient %s: %s", recipient.Hex(), amount.String())
			return &AirdropError{
				Code:    "INVALID_AMOUNT",
				Message: fmt.Sprintf("Invalid amount for recipient %s", recipient.Hex()),
			}
		}
	}

	log.Println("[SUCCESS] Airdrop allocations validated successfully.")
	return nil
}

// CreateAndSendBundledTx sends multiple token transfers and returns a map of recipient addresses to transaction hashes
func (a *Airdrop) CreateAndSendBundledTx(allocations map[common.Address]*big.Int) (map[common.Address]common.Hash, error) {
	log.Println("[PROCESS] Starting bundled transactions for airdrop...")

	// Transaction hash map to store results
	txHashes := make(map[common.Address]common.Hash)

	if err := a.ValidateAirdropAllocation(allocations); err != nil {
		log.Printf("[ERROR] Validation failed: %v", err)
		return txHashes, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(a.PrivateKey, a.ChainID)
	if err != nil {
		log.Printf("[ERROR] Failed to create transaction options: %v", err)
		return txHashes, &AirdropError{
			Code:    "TRANSACTOR_CREATION_FAILED",
			Message: "Failed to create transaction options",
			Err:     err,
		}
	}

	var failedTransfers []string
	for recipient, amount := range allocations {
		if recipient == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			log.Printf("[WARN] Skipping transfer to zero address")
			continue
		}

		log.Printf("[TRANSFER] Sending %s tokens to %s...", amount.String(), recipient.Hex())

		tx, err := a.Token.Transfer(opts, recipient, amount)
		if err != nil {
			failedTransfers = append(failedTransfers, fmt.Sprintf("Recipient: %s | Error: %v", recipient.Hex(), err))
			log.Printf("[ERROR] Transfer failed for recipient %s: %v", recipient.Hex(), err)
			continue
		}

		// Store transaction hash for the successful transfer
		txHashes[recipient] = tx.Hash()
		log.Printf("[SUCCESS] Transaction submitted for %s: Tx Hash: %s", recipient.Hex(), tx.Hash().Hex())
	}

	if len(failedTransfers) > 0 {
		log.Printf("[WARNING] Some transfers failed:\n%s", strings.Join(failedTransfers, "\n"))
		return txHashes, &AirdropError{
			Code:    "PARTIAL_TRANSFER_FAILURE",
			Message: strings.Join(failedTransfers, "\n"),
		}
	}

	log.Println("[SUCCESS] All transfers completed successfully.")
	return txHashes, nil
}

//func main() {
//	rpcURL := "https://arbitrum-mainnet.infura.io/v3/23c96151438348e79d1e963324d33b02"
//	privateKeyHex := "f0029d055d03c9e06e4c87ac6f0c6209948707987765d47e6121e873d9fae3cf"
//	tokenAddressHex := "0xaf88d065e77c8cc2239327c5edb3a432268e5831"
//
//	airdrop, err := NewAirdrop(rpcURL, privateKeyHex, tokenAddressHex)
//	if err != nil {
//		log.Fatalf("[FATAL] Failed to initialize airdrop: %v", err)
//	}
//	defer airdrop.Client.Close()
//
//	allocations := map[common.Address]*big.Int{
//		common.HexToAddress("0xCc52829336f610B6f22A8C381Fb8Dd05a7BDba93"): MustConvertToTokenUnits(airdrop.Token, 0.01),
//		common.HexToAddress("0x034a084D73e5FC65308394B45A3235BA48a623ac"): MustConvertToTokenUnits(airdrop.Token, 0.01),
//		common.HexToAddress("0xE164004bB260EC5c4E007837DA21Be984419054f"): MustConvertToTokenUnits(airdrop.Token, 0.01),
//	}
//
//	txHashes, err := airdrop.CreateAndSendBundledTx(allocations)
//	if err != nil {
//		log.Fatalf("[FATAL] Failed to complete airdrop transactions: %v", err)
//	}
//
//	// Optional: Print out transaction hashes
//	for recipient, txHash := range txHashes {
//		log.Printf("Recipient %s received tokens in transaction: %s", recipient.Hex(), txHash.Hex())
//	}
//
//	log.Println("[COMPLETE] Airdrop process completed successfully.")
//}
