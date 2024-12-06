package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// SolanaAirdropError provides structured error handling similar to the Ethereum implementation
type SolanaAirdropError struct {
	Code    string
	Message string
	Err     error
}

func (e *SolanaAirdropError) Error() string {
	log.Printf("[ERROR] Code: %s | Message: %s | Details: %v", e.Code, e.Message, e.Err)
	return fmt.Sprintf("%s: %s (underlying error: %v)", e.Code, e.Message, e.Err)
}

// SolanaAirdrop struct to encapsulate airdrop functionality
type SolanaAirdrop struct {
	Client     *client.Client
	FeePayer   types.Account
	MintPubkey common.PublicKey
}

// NewSolanaAirdrop creates a new SolanaAirdrop instance
func NewSolanaAirdrop(rpcEndpoint, feePayer, mintPubkeyStr string) (*SolanaAirdrop, error) {
	log.Printf("[INIT] Creating SolanaAirdrop instance with RPC Endpoint: %s, Mint Address: %s", rpcEndpoint, mintPubkeyStr)

	if rpcEndpoint == "" || feePayer == "" || mintPubkeyStr == "" {
		log.Printf("[ERROR] Missing required parameters: RPC endpoint, fee payer, or mint address.")
		return nil, &SolanaAirdropError{
			Code:    "INVALID_PARAMS",
			Message: "RPC endpoint, fee payer, and mint address cannot be empty",
		}
	}

	log.Println("[CONNECT] Establishing connection to Solana client...")
	client := client.NewClient(rpcEndpoint)

	// Convert fee payer from base58
	feePayerAccount, err := types.AccountFromBase58(feePayer)
	if err != nil {
		log.Printf("[ERROR] Invalid fee payer account: %v", err)
		return nil, &SolanaAirdropError{
			Code:    "INVALID_FEE_PAYER",
			Message: "Failed to parse fee payer account",
			Err:     err,
		}
	}

	// Convert mint pubkey
	mintPubkey := common.PublicKeyFromString(mintPubkeyStr)

	log.Println("[SUCCESS] SolanaAirdrop instance created successfully.")
	return &SolanaAirdrop{
		Client:     client,
		FeePayer:   feePayerAccount,
		MintPubkey: mintPubkey,
	}, nil
}

func (a *SolanaAirdrop) ValidateAirdropAllocation(allocations map[string]float64) error {
	log.Println("[VALIDATION] Validating airdrop allocations...")

	totalAirdropAmount := 0.0
	for recipient, amount := range allocations {
		log.Printf("[ALLOCATION] Recipient: %s | Amount: %f", recipient, amount)

		if amount <= 0 {
			log.Printf("[ERROR] Invalid allocation amount for recipient %s: %f", recipient, amount)
			return &SolanaAirdropError{
				Code:    "INVALID_AMOUNT",
				Message: fmt.Sprintf("Invalid amount for recipient %s", recipient),
			}
		}

		totalAirdropAmount += amount
	}

	log.Printf("[TOTAL] Total airdrop amount: %f", totalAirdropAmount)
	log.Println("[SUCCESS] Airdrop allocations validated successfully.")
	return nil
}

func (a *SolanaAirdrop) CreateAndSendBundledTx(allocations map[string]float64) (map[string]string, error) {
	log.Println("[PROCESS] Starting bundled transactions for airdrop...")

	txHashes := make(map[string]string)

	if err := a.ValidateAirdropAllocation(allocations); err != nil {
		log.Printf("[ERROR] Validation failed: %v", err)
		return txHashes, err
	}

	res, err := a.Client.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Printf("[ERROR] Failed to get latest blockhash: %v", err)
		return txHashes, &SolanaAirdropError{
			Code:    "BLOCKHASH_FETCH_FAILED",
			Message: "Failed to fetch latest blockhash",
			Err:     err,
		}
	}

	decimals, err := a.GetDecimals(a.Client, a.MintPubkey, a.FeePayer.PublicKey)

	var failedTransfers []string
	for recipient, amount := range allocations {
		log.Printf("[TRANSFER] Sending %f tokens to %s...", amount, recipient)

		// Derive recipient's Associated Token Account (ATA)
		recipientPubkey := common.PublicKeyFromString(recipient)
		recipientATA, _, err := common.FindAssociatedTokenAddress(recipientPubkey, a.MintPubkey)
		if err != nil {
			errMsg := fmt.Sprintf("Recipient: %s | Error: Failed to derive ATA: %v", recipient, err)
			failedTransfers = append(failedTransfers, errMsg)
			log.Printf("[ERROR] %s", errMsg)
			continue
		}
		feePayerATA, _, err := common.FindAssociatedTokenAddress(a.FeePayer.PublicKey, a.MintPubkey)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to derive Fee Payer ATA: %v", err)
			failedTransfers = append(failedTransfers, errMsg)
			log.Printf("[ERROR] %s", errMsg)
			continue
		}

		// Build transaction
		tx, err := types.NewTransaction(types.NewTransactionParam{
			Message: types.NewMessage(types.NewMessageParam{
				FeePayer:        a.FeePayer.PublicKey,
				RecentBlockhash: res.Blockhash,
				Instructions: []types.Instruction{
					token.TransferChecked(token.TransferCheckedParam{
						From:     feePayerATA,
						To:       recipientATA,
						Mint:     a.MintPubkey,
						Auth:     a.FeePayer.PublicKey,
						Signers:  []common.PublicKey{},
						Amount:   uint64(amount),
						Decimals: decimals, // Update based on actual token decimals
					}),
				},
			}),
			Signers: []types.Account{a.FeePayer},
		})
		if err != nil {
			errMsg := fmt.Sprintf("Recipient: %s | Error: Failed to create transaction: %v", recipient, err)
			failedTransfers = append(failedTransfers, errMsg)
			log.Printf("[ERROR] %s", errMsg)
			continue
		}

		// Send transaction
		txHash, err := a.Client.SendTransaction(context.Background(), tx)
		if err != nil {
			errMsg := fmt.Sprintf("Recipient: %s | Error: Failed to send transaction: %v", recipient, err)
			failedTransfers = append(failedTransfers, errMsg)
			log.Printf("[ERROR] %s", errMsg)
			continue
		}

		// Store successful transaction hash
		txHashes[recipient] = txHash
		log.Printf("[SUCCESS] Transaction submitted for %s: Tx Hash: %s", recipient, txHash)
	}

	if len(failedTransfers) > 0 {
		log.Printf("[WARNING] Some transfers failed:\n%s", strings.Join(failedTransfers, "\n"))
		return txHashes, &SolanaAirdropError{
			Code:    "PARTIAL_TRANSFER_FAILURE",
			Message: strings.Join(failedTransfers, "\n"),
		}
	}

	log.Println("[SUCCESS] All transfers completed successfully.")

	return txHashes, nil
}

func (a *SolanaAirdrop) GetDecimals(c *client.Client, mintAddress common.PublicKey, userWalletAddress common.PublicKey) (uint8, error) {

	// Derive the associated token account address
	ata, _, err := common.FindAssociatedTokenAddress(userWalletAddress, mintAddress)\
	if err != nil {
		return 0, fmt.Errorf("failed to derive associated token account: %w", err)
	}

	// Fetch the token account balance and decimals
	tokenBalance, err := c.GetTokenAccountBalance(context.Background(), ata.ToBase58())
	if err != nil {
		return 0, fmt.Errorf("failed to fetch token decimals: %v", err)
	}
	decimals := tokenBalance.Decimals

	return decimals, nil
}

func main() {
	rpcURL := rpc.DevnetRPCEndpoint
	feePayer := "Hd8j2ffdBnV6deHRyy68RCa1jDsH6VWyjG6XhJe91YibZ32yLc1w9jskia3bv4bBbgC3pmMJaeajQMUEztaHCeH"
	mintPubkeyStr := "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"

	airdrop, err := NewSolanaAirdrop(rpcURL, feePayer, mintPubkeyStr)
	if err != nil {
		log.Fatalf("[FATAL] Failed to initialize airdrop: %v", err)
	}

	allocations := map[string]float64{
		"G1UgQHs2z8Vh6m7QwTVPsiKsgUUQLH5g2SuVeJoXpWu":  0.1,
		"EeFd1G5HwtXUtuZj7Jsm3y6jPL4h6QoSy44Bsp2HV8aU": 0.1,
		"3baWoQybGod1X7LFEeErCZCbNvSQN8jaCZv64Cuic22H": 0.1,
	}

	// Execute airdrop
	txHashes, err := airdrop.CreateAndSendBundledTx(allocations)
	if err != nil {
		log.Fatalf("[FATAL] Failed to complete airdrop transactions: %v", err)
	}

	// Print transaction hashes
	for recipient, txHash := range txHashes {
		log.Printf("Recipient %s received tokens in transaction: %s", recipient, txHash)
	}

	log.Println("[COMPLETE] Airdrop process completed successfully.")
}
