package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list you local installed dotnet core sdk version",
	Long: `examples of using this command. For example:
    dvm list   -- list all local sdk`,
	Run: func(cmd *cobra.Command, args []string) {
		getLocalList()
	},
}

func getLocalList() {
	sdksPath := getDvmSdkStorePath("")
	cmd := exec.Command("ls", sdksPath)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("list installed sdk failed. ", contactUs)
	}
	fmt.Println(string(out))
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
