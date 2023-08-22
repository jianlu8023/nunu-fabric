package chaincode

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	gmgw "github.com/hxx258456/fabric-sdk-go-gm/pkg/gateway"
)

func gmCreateWallet(wallet *gmgw.Wallet) error {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(fmt.Errorf("get current work dir has error: %s", err))
	}
	credPath := filepath.Join(wd, "..", "testdata", "fixtures", "organizations", "peerOrganizations", "org1.example.com", "users", "User1@org1.example.com", "msp")

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")

	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")

	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}
	identity := gmgw.NewX509Identity("Org1MSP", string(cert), string(key))
	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
