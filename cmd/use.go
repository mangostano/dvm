package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "change local dotnet core version",
	Long: `examples of using this command. For example: 
   dvm use 1.1  -- this will use the LTS version 
   dvm use 1.1.2 -- this will use the specific version 1.1.2
`,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		if !checkSdkMainVersionExists(version) {
			if checkSdkSubVersionExists(version) {
				version = latestSubVersion(version)
			} else {
				installCmd := exec.Command("dvm", "install", version)
				if err := installCmd.Run(); err != nil {
					log.Fatal("please use dvm install ", version, " to install this command")
				}

			}
		}
		removeOtherLink()
		createLink(version)
		fmt.Println("now you dotnet core sdk version is: ", version)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
