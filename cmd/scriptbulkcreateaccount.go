package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
 	"strconv"
)

var scriptBulkCreateAccount = &cobra.Command{
	Use:   "bulkaccount",
	Short: "bulk create validators and execute deposit tx on eth 1.0",
	Long: `Create an account.  For example:

    ethdo script bulkaccount --wallet="primary" --walletpassphrase="my secret"

	`,
	Run: func(cmd *cobra.Command, args []string) {
		assert(walletName != "", "--wallet is required")
		assert(bulkNumber != "", "--bulk is required")
		assert(rootWalletPassphrase != "", "--passphrase is required")

		cnt, err := strconv.Atoi(bulkNumber)
		errCheck(err, "--bulkNumber invalid value")

		for i := 1;  i<=cnt; i++ {
			accountRoot := walletName + "/" + strconv.Itoa(i)

			w, err := walletFromPath(accountRoot)
			errCheck(err, "Failed to access wallet")

			_, err = accountFromPath(accountRoot)
			assert(err != nil, "Account already exists")

			err = w.Unlock([]byte(rootWalletPassphrase))
			errCheck(err, "Failed to unlock wallet")

			_, accountName, err := walletAndAccountNamesFromPath(accountRoot)
			errCheck(err, "Failed to obtain accout name")

			account, err := w.CreateAccount(accountName, []byte(rootAccountPassphrase))
			errCheck(err, "Failed to create account")

			outputIf(verbose, fmt.Sprintf("0x%048x", account.PublicKey().Marshal()))
		}
	},
}

func init() {
	scriptCmd.AddCommand(scriptBulkCreateAccount)
	scriptFlags(scriptBulkCreateAccount)
}