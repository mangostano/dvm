package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use: "install",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 0 {
			errors.New("not have a valid args found, please use 'dvm install -h' to get more info")
		}
		return nil
	},
	Short: "This command use to install sdk",
	Long: `This is simply brief introduce the 'install' command 
	here is the basic usage
	'dvm install 1.1' this will install dotnet core 1.1 LTS `,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		version = download(version)
		moveDotnetVersion(version)
		removeOtherLink()
		createLink(version)
		fmt.Println("install completely")
		fmt.Println("now you dotnet core version is ", version, "please enjoy")
	},
}

func download(version string) string {
	if checkSdkMainVersionExists(version) {
		return version
	}
	if checkSdkSubVersionExists(version) {
		return latestSubVersion(version)
	}
	dvmHome := getDvmHome()
	installFile := fmt.Sprint(dvmHome, "/scripts", "/install.sh")
	args := []string{"-v", version}
	cmd := exec.Command(installFile, args...)
	fmt.Println("starting to install dotnet core sdk, this will take a few minutes, please wait!")
	err := cmd.Run()
	if err != nil {
		subVersion := latestSubVersion(version)
		retryCmd := exec.Command(installFile, "-v", subVersion)
		output, error := retryCmd.Output()
		if error != nil {
			log.Fatal("unknown dotnet version please use `dvm listAll` to check the install version is correct")
		}
		fmt.Println(output)
		return subVersion
	}
	return version
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
