package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "This command use to uninstall sdk",
	Args:  cobra.NoArgs,
	Long: `Simple usage:
 dvm current'`,
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion := strings.TrimSpace(getUsingVersion())
		fmt.Println(currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
