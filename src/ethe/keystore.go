package ethe

import (
	k "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

const WalletPath string = "./wallet/"

func NewAccount(pwd string) (accounts.Account, error) {
	a := accounts.Account{}
	ks := k.NewKeyStore(WalletPath, 2, 1) // Todo: scryptN, scryptP
	a, err := ks.NewAccount(pwd)
	if err != nil {
		return a, err
	}
	return a, nil
}

type KeystoreV3 struct {
	json encryptedKeyJSONV3
	pwd  string
}

type encryptedKeyJSONV3 struct {
	Address string       `json:"address"`
	Crypto  k.CryptoJSON `json:"crypto"`
	Id      string       `json:"id"`
	Version int          `json:"version"`
}

func readKeystore(path string, pwd string) KeystoreV3 {
	en := encryptedKeyJSONV3{}
	jsonFile, _ := os.Open(path)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &en)

	return KeystoreV3{json: en, pwd: pwd}

}

func (key KeystoreV3) Address() common.Address {
	return common.HexToAddress(key.json.Address)
}

func (key KeystoreV3) PrivateKey() ([]byte, error) {
	//privBytes, err := k.DecryptDataV3(key.Crypto, pwd)
	//hex.EncodeToString(privBytes)
	return k.DecryptDataV3(key.json.Crypto, key.pwd)
}

func ResultKeystore(index int, pwd string) (KeystoreV3, error) {
	keyV3 := KeystoreV3{}

	files, err := ioutil.ReadDir(WalletPath)
	if err != nil {
		log.Fatal(err)
	}

	if (index + 1) > len(files) {
		log.Println("index out of range")
		return keyV3, err
	}

	for i, file := range files {
		if i == index {
			keyV3 = readKeystore(WalletPath+file.Name(), pwd)
			break
		}
	}

	return keyV3, nil
}
