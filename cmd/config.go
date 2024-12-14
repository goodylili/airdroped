package configurations

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	RPCURL        string
	PrivateKey    string
	TokenAddress  string
	CSVFilePath   string
	AddressColumn string
	AmountsColumn string
	AirdropOption string
}

func LoadConfigurations() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, proceeding with system environment variables")
	}

	// Populate and return the configuration struct
	return &Config{
		RPCURL:        os.Getenv("RPC_URL"),
		PrivateKey:    os.Getenv("PRIVATE_KEY"),
		TokenAddress:  os.Getenv("TOKEN_ADDRESS"),
		CSVFilePath:   os.Getenv("CSV_FILE_PATH"),
		AddressColumn: os.Getenv("ADDRESS_COLUMN"),
		AmountsColumn: os.Getenv("AMOUNTS_COLUMN"),
		AirdropOption: os.Getenv("AIRDROP_OPTION"),
	}
}
