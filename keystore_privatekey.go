package main

import (
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"fmt"
	"github.com/urfave/cli"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func KeysotreToPrivatekey(filename, auth string) (*ecdsa.PrivateKey, common.Address, error) {
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, common.Address{}, err
	}

	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, common.Address{}, err
	}

	return key.PrivateKey, key.Address, nil
}

func main() {

	app := cli.NewApp()
	app.Name = "keystore-privatekey"
	app.Usage = "Extract private key from keystore file with passphrase"
	app.Version = "0.0.1"
	app.UsageText = "keystore-privatekey [keystore path]"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "toconsole",
			Usage: "output to console",
		},
	}

	app.Action = func(c *cli.Context) error {
		// 1. Get keystore path
		filepath := c.Args().Get(0)
		if len(filepath) == 0 {
			return errors.New("need pass keystore path as first argument \nusage: keystore_privatekey [keystore path]")
		}

		// 2. Get passphrase
		auth, err := console.Stdin.PromptPassword("passphrase: ")

		// 3. Generate private key
		privkey, addr, err := KeysotreToPrivatekey(filepath, auth)
		if err != nil {
			return err
		}

		// 4. Transfer to hex
		privkeyHex := hex.EncodeToString(crypto.FromECDSA(privkey))

		// 5. write to destination
		toconsole := c.BoolT("toconsole")
		if !toconsole {
			// write to {address}.txt
			outputName := addr.String() + ".txt"
			if err := ioutil.WriteFile(outputName, []byte(privkeyHex), 0600); err != nil {
				return errors.New("cannot write to " + outputName)
			}
			fmt.Println("have written to file: " + outputName)

		} else {
			fmt.Println(privkeyHex)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}
