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
			fmt.Println("[Info] now your are download the latest DVM version, this will be take a few minutes, please wait...")
			removeOldDvmVersion()
			installLatestDvmVersion(latestDvmVersion)
			fmt.Println("[Info] DVM upgrade finish, now your DVM version is ", latestDvmVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
