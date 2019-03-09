package cmd

import (
	"fmt"
    "os/exec"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "This command use to install sdk",
	Long: `This is simply brief introduce the 'install' command 
	here is the basic usage
	'dvm install 1.1' this will install dotnet core 1.1 sdk `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
		 fmt.Println(`not have a valid args found, please use 'dvm install --help' to get more info`);
		 return ;
		}
		fmt.Println("install called")
	},
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
