package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "This command use to uninstall sdk",
	Long: `This is simply brief introduce the 'uninstall' command 
	here is the basic usage
	'dvm uninstall 1.1' this will uninstall dotnet core 1.1 LTS`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) > 1 {
			fmt.Println(`not have a valid args found, please use 'dvm uninstall --help' to get more info`)
			return
		}

		version := strings.TrimSpace(args[0])
		installVersion := getInstallVersions()

		if !contains(installVersion, version) {
			log.Fatal("The uninstall version:" + version + " not in location")
		}

		if version == strings.TrimSpace(getUsingVersion()) {
			log.Fatal("The uninstall version:" + version + " is using, please change using version")
		}

		err := deleteSDK(version)
		if err != nil {
			log.Fatal("The input version:" + version + " is a wrong version or not installed")
		}
		fmt.Println("uninstall completely")
	},
}

func deleteSDK(version string) error {
	args := []string{"-rf", getDvmSdkStorePath(version)}
	cmd := exec.Command("rm", args...)

	err := cmd.Run()
	return err
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
