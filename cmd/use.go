package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "change local dotnet core version",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use called")
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
