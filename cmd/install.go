package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const contactUs = "please contact us.(https://github.com/mangostano/dvm/issues)"

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "This command use to install sdk",
	Long: `This is simply brief introduce the 'install' command 
	here is the basic usage
	'dvm install 1.1' this will install dotnet core 1.1 LTS `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`not have a valid args found, please use 'dvm install --help' to get more info`)
			return
		}
		version := args[0]
		download(version)
		moveDotnetVersion(version)
		removeOtherLink()
		createLink(version)
		fmt.Println("install completely")
		fmt.Println("now you dotnet core version is ", version, "please enjoy")
	},
}

func removeOtherLink() {
	args := []string{"-rf", fmt.Sprint(getDotnetHome(), "/sdk/*")}
	cmd := exec.Command("rm", args...)
	err := cmd.Run()
	if err != nil {
		log.Fatal("install failed when remove other link. ", contactUs)
	}
}

func createLink(version string) {
	source := getDvmSdkStorePath(version)
	dest := getDotnetSdkPath("")
	cmd := exec.Command("ln", "-s", source, dest)
	err := cmd.Run()
	if err != nil {
		log.Fatal("install failed when crate link. ", contactUs)
	}
}

func moveDotnetVersion(version string) {
	descPath := getDvmSdkStorePath(version)
	sourcePath := getDotnetSdkPath(version)
	moveFile(sourcePath, descPath)
}

func getDotnetSdkPath(version string) string {
	if len(version) <= 0 {
		return fmt.Sprint(getDotnetHome(), "/sdk/")
	}
	return fmt.Sprint(getDotnetHome(), "/sdk/", version)
}

func getDvmSdkStorePath(version string) string {
	if len(version) <= 0 {
		return fmt.Sprint(getDvmHome(), "/sdks")
	}
	return fmt.Sprint(getDvmHome(), "/sdks/", version)
}

func moveFile(sourceDir string, destDir string) {
	args := []string{"-f", sourceDir, destDir}
	cmd := exec.Command("mv", args...)
	err := cmd.Run()
	if err != nil {
		retryCmd := exec.Command("mv", "-rf", fmt.Sprint(sourceDir, "*"), destDir)
		retryError := retryCmd.Run()
		if retryError != nil {
			log.Fatal("install failed when move file. ", contactUs)
		}
	}
}

func download(version string) {
	dvmHome := getDvmHome()
	installFile := fmt.Sprint(dvmHome, "/scripts", "/install.sh")
	args := []string{"-v", version}
	cmd := exec.Command(installFile, args...)
	fmt.Println("starting to install dotnet core sdk, this will take a few minutes, please wait!")
	err := cmd.Run()
	if err != nil {
		retryCmd := exec.Command(installFile, "-c", "version")
		error := retryCmd.Run()
		if error != nil {
			log.Fatal("unknown dotnet version please use `dvm ls` to check the install version is correct")
			return
		}
	}
}

func getDvmHome() string {
	dvmHome := os.Getenv("DVM_HOME")
	if len(dvmHome) <= 0 {
		dvmHome = fmt.Sprint(os.Getenv("HOME"), "/.dvm")
	}
	return dvmHome
}

func getDotnetHome() string {

	dotnetHome := os.Getenv("DOTNET_HOME")
	if len(dotnetHome) <= 0 {
		dotnetHome = fmt.Sprint(os.Getenv("HOME"), "/.dotnet")
	}
	return dotnetHome
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
