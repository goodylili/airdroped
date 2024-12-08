package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
)

type AirdropConfig struct {
	Mnemonic string
	ObjectID string // The object ID of the token to airdrop
}

type AirdropResult struct {
	Address  string
	Amount   string
	Success  bool
	TxDigest string
	GasUsed  string
	Error    error
}

func ProcessAirdrop(ctx context.Context, config AirdropConfig, recipients map[string]string) []AirdropResult {
	cli := sui.NewSuiClient(constant.BvMainnetEndpoint)

	signerAccount, err := signer.NewSignertWithMnemonic(config.Mnemonic)
	if err != nil {
		log.Fatalf("Failed to create signer: %v", err)
	}

	results := make([]AirdropResult, 0, len(recipients))

	for address, amount := range recipients {
		result := AirdropResult{
			Address: address,
			Amount:  amount,
		}

		// Prepare transfer transaction
		txnBlock, err := cli.TransferObject(ctx, models.TransferObjectRequest{
			Signer:    signerAccount.Address,
			ObjectId:  config.ObjectID,
			Recipient: address,
		})

		if err != nil {
			result.Error = fmt.Errorf("failed to prepare transfer: %v", err)
			result.Success = false
			results = append(results, result)
			continue
		}

		// Simulate transaction to get gas estimate
		inspectResp, err := cli.DevInspectTransactionBlock(ctx, models.DevInspectTransactionBlockRequest{
			TxBytes:    txnBlock,
			Sender:     signerAccount.Address,
			GasPrice:   "1000",
			SkipChecks: true,
		})

		if err != nil {
			result.Error = fmt.Errorf("failed to estimate gas: %v", err)
			result.Success = false
			results = append(results, result)
			continue
		}

		// Add gas budget based on simulation
		txnBlock.GasBudget = inspectResp.Effects.GasUsed
		result.GasUsed = inspectResp.Effects.GasUsed

		// Sign and execute the transaction
		txnResponse, err := cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
			TxnMetaData: txnBlock,
			PriKey:      signerAccount.PriKey,
			Options: models.SuiTransactionBlockOptions{
				ShowEffects: true,
			},
			RequestType: "WaitForLocalExecution",
		})

		if err != nil {
			result.Error = fmt.Errorf("failed to execute transfer: %v", err)
			result.Success = false
		} else {
			result.Success = true
			result.TxDigest = txnResponse.Digest
		}

		results = append(results, result)
	}

	return results
}

func main() {
	config := AirdropConfig{
		Mnemonic: "your twelve word mnemonic phrase here",
		ObjectID: "token-object-id-to-airdrop",
	}

	recipients := map[string]string{
		"0xaddress1": "100",
		"0xaddress2": "200",
		"0xaddress3": "300",
	}

	ctx := context.Background()
	results := ProcessAirdrop(ctx, config, recipients)

	for _, result := range results {
		if result.Success {
			fmt.Printf("Successfully sent %s to %s (TX: %s, Gas Used: %s)\n",
				result.Amount, result.Address, result.TxDigest, result.GasUsed)
		} else {
			fmt.Printf("Failed to send %s to %s: %v\n",
				result.Amount, result.Address, result.Error)
		}
	}
}
func main() {
	// Example usage
	config := AirdropConfig{
		Mnemonic:    "energy make under forward clip dose congress wait salad expect betray initial",
		GasObjectID: "0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI",
		GasBudget:   "100000000",
		ObjectID:    "token-object-id-to-airdrop",
	}

	// Create a map of recipients and their amounts
	recipients := map[string]string{
		"0xaddress1": "100",
		"0xaddress2": "200",
		"0xaddress3": "300",
	}

	ctx := context.Background()
	results := ProcessAirdrop(ctx, config, recipients)

	// Print results
	for _, result := range results {
		if result.Success {
			fmt.Printf("Successfully sent %s to %s (TX: %s)\n",
				result.Amount, result.Address, result.TxDigest)
		} else {
			fmt.Printf("Failed to send %s to %s: %v\n",
				result.Amount, result.Address, result.Error)
		}
	}
}
