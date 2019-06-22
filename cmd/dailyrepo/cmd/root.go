package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:   "dailyrepo",
	Short: "日報作成ツール",
	Long: `テンプレートから日報の雛形を作成します`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			printVersion()
			return nil
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "print version")
	rootCmd.Flags().Bool("verbose", false, "print log")
}

func printVersion() {
	fmt.Printf("dailyrepo %s\n", version)
}
