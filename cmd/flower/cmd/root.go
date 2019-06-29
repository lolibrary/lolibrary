package cmd

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/spf13/cobra"
)

var (
	env           string
	createForward bool
)

var rootCmd = &cobra.Command{
	Use:   "flower",
	Short: "🌸 Lolibrary's interactive RPC shell",
	Long: `🌸 Lolibrary's interactive RPC shell.

	Provides line history and pretty-printed output,
	as well as port-forwarding to the edge-proxy automatically.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if createForward {
			port, forward := portforward.Enable()
			defer forward.Close()

			EnvProd = fmt.Sprintf("http://localhost:%d", port)
		}

		if err := environment(); err != nil {
			return err
		}

		shellCmd.Run(cmd, args)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "local",
		fmt.Sprintf("set the environment used (valid: %v, %v)", aurora.Green("prod"), aurora.Green("local")))
	rootCmd.PersistentFlags().BoolVarP(&createForward, "forward", "f", false, "enable automatic port-forwarding to the edge proxy")
	rootCmd.AddCommand(shellCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
