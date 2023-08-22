package chaincode

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	gmconfig "github.com/hxx258456/fabric-sdk-go-gm/pkg/core/config"
	gmgw "github.com/hxx258456/fabric-sdk-go-gm/pkg/gateway"
)

func Test_gm_basic_QueryAllAssert(t *testing.T) {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(fmt.Errorf("get current work dir has error: %s", err))
	}
	wallet, err := gmgw.NewFileSystemWallet(filepath.Clean(filepath.Join(wd, "..", "testdata", "wallet")))
	if err != nil {
		fmt.Println(fmt.Errorf("get file system wallet has error: %s", err))
	}
	if wallet.Exists("appUser") {
		err = gmCreateWallet(wallet)
		if err != nil {
			fmt.Println(fmt.Sprintf("create user appUser has error:%s", err))
		}
	}

	connectionPath := filepath.Join(wd, "..", "testdata", "fixtures", "organizations", "peerOrganizations", "org1.example.com", "connection-org1.yaml")

	gw, err := gmgw.Connect(
		gmgw.WithConfig(gmconfig.FromFile(filepath.Clean(connectionPath))),
		gmgw.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("get gatewayClient has error:%s", err))

	}

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Println(fmt.Sprintf("get channel mychannel has error:%s", err))

	}

	contract := network.GetContract("baisc")
	allAssert, err := contract.EvaluateTransaction("GetAllAssets")

	if err != nil {
		fmt.Println(fmt.Sprintf("call GetAllAsserts has error:%s", err))
	}

	fmt.Println(string(allAssert))
}

func Test_getWD(t *testing.T) {
	// method 1
	wd, _ := os.Getwd()
	fmt.Println("os.Getwd(): ", wd)

	lookPath, _ := exec.LookPath(os.Args[0])
	abs, _ := filepath.Abs(lookPath)
	dir := path.Dir(abs)
	fmt.Println("os.Args[0]: ", dir)
	_, file, _, _ := runtime.Caller(0)
	fmt.Println("runtime.Caller(0): ", file)

}
