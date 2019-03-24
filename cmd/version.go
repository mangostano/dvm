package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.NoArgs,
	Short: "This command to show the DVM version",
	Long: `Simple usage:  
 dvm version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your current DVM version is :", currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
