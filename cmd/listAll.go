package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "This command to get all of the dotnet core sdk versions",
	Long: `examples of using this command. For example:
		dvm listAll`,
	Run: func(cmd *cobra.Command, args []string) {
		result := make(map[string][]string)
		getVersionJsonFile(fmt.Sprintf(versionFileUrlTemplate, currentVersion), result)
		printResult(result)
	},
}

func printResult(result map[string][]string) {
	var mainVersions []string
	for mainVersion := range result {
		mainVersions = append(mainVersions, mainVersion)
	}
	sort.Strings(mainVersions)
	for _, mainVersion := range mainVersions {
		fmt.Println("    ", mainVersion)
		for _, subVersion := range result[mainVersion] {
			fmt.Println("\t", subVersion)
		}
	}
}

func init() {
	rootCmd.AddCommand(listAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
