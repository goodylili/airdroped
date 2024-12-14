package configurations

import (
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

type Config struct {
	RPCURL        string
	PrivateKey    string
	TokenAddress  string
	CSVFilePath   string
	AddressColumn string
	AmountsColumn string
	ChainID       *big.Int
	AirdropOption string
	CA            string
	PublicKey     string
}

func LoadConfigurations() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, proceeding with system environment variables")
	}
	chainIDStr := os.Getenv("CHAIN_ID")
	chainID := new(big.Int)
	if chainIDStr != "" {
		_, success := chainID.SetString(chainIDStr, 10) // Base 10 for decimal parsing
		if !success {
			log.Fatalf("Invalid CHAIN_ID: %s", chainIDStr)
		}
	} else {
		log.Fatalf("CHAIN_ID is required but not set")
	}
	return &Config{
		RPCURL:        os.Getenv("RPC_URL"),
		PrivateKey:    os.Getenv("PRIVATE_KEY"),
		TokenAddress:  os.Getenv("TOKEN_ADDRESS"),
		ChainID:       chainID,
		PublicKey:     os.Getenv("PUBLIC_KEY"),
		CA:            os.Getenv("CONTRACT_DEPLOYMENT_ADDRESS"),
		CSVFilePath:   os.Getenv("CSV_FILE_PATH"),
		AddressColumn: os.Getenv("ADDRESS_COLUMN"),
		AmountsColumn: os.Getenv("AMOUNTS_COLUMN"),
		AirdropOption: os.Getenv("AIRDROP_OPTION"),
	}
}
