package evm

//// CreateWallet generates a new Ethereum wallet for user.
//func CreateWallet() (address, privateKey string, err error) {
//	privateKeyECDSA, err := crypto.GenerateKey()
//	if err != nil {
//		return "", "", err
//	}
//	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
//	privateKey = hexutil.Encode(privateKeyBytes)[2:]
//
//	publicKey := privateKeyECDSA.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("error casting public key to ECDSA")
//	}
//
//	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
//
//	return address, privateKey, nil
//}
