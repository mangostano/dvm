package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your dvm version",
	Run: func(cmd *cobra.Command, args []string) {
		latestDvmVersion := getLatestDvmVersion()
		if strings.EqualFold(latestDvmVersion, currentVersion) {
			fmt.Println("[INFO] Your DVM version is the latest version, no need upgrade!")
		} else {
			removeOldDvmVersion()
			installLatestDvmVersion(latestDvmVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
