package sui

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenerateSuiWallet() (string, string, error) {
	cmd := exec.Command("sui", "keytool", "generate", "ed25519", "--json")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	var keyInfo KeyInfo
	err = json.Unmarshal(out.Bytes(), &keyInfo)
	if err != nil {
		return "", "", err
	}
	dir, _ := os.Getwd()
	files, _ := filepath.Glob(filepath.Join(dir, "*.key"))
	for _, file := range files {
		os.Remove(file)
	}
	return keyInfo.SuiAddress, strings.TrimSpace(keyInfo.Mnemonic), nil

}
