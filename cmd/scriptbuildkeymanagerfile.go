package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	types "github.com/wealdtech/go-eth2-wallet-types"
	"log"
)

type KeyManagerData struct {
	Location    string `json:"location"`
	Accounts    []string `json:"accounts"`
	Passphrases []string `json:"passphrases"`
}

var scriptBuildKeyManagerFile = &cobra.Command{
	Use:   "keymanagerfile",
	Short: "generate a keymanager file for prysm validator client",
	Long: `For example:

    ethdo script keymanagerfile --wallet="primary" --passphrase="my secret"

	`,
	Run: func(cmd *cobra.Command, args []string) {
		assert(walletName != "", "--wallet is required")
		//assert(rootAccountPassphrase != "", "--passphrase is required")

		// get accounts
		wallet, err := walletFromPath(walletName)
		errCheck(err, "Failed to access wallet")

		// get all accounts
		accounts := []types.Account{<- wallet.Accounts()}
		for account := range wallet.Accounts() {
			accounts = append(accounts, account)
		}

		// prepare passphrases array
		_pass := make([]string, len(accounts))
		for i := 1;  i<len(accounts); i++ {
			_pass[i] = rootAccountPassphrase
		}

		data := KeyManagerData{
			Location:    "/wallets",
			Accounts:    []string{walletName + "/*"},
			Passphrases: _pass,
		}

		jsonData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(jsonData))
	},
}

func init() {
	scriptCmd.AddCommand(scriptBuildKeyManagerFile)
	scriptFlags(scriptBuildKeyManagerFile)
}