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
		dvm listAll     --This will be list all versions
        dvm listAll 1.1  --This will be list 1.1's sub versions`,
	Run: func(cmd *cobra.Command, args []string) {
		result := make(map[string][]string)
		getVersionJsonFile(fmt.Sprintf(versionFileUrlTemplate, currentVersion), result)
		printResult(result, args)
	},
}

func printResult(result map[string][]string, args []string) {
	var mainVersions []string
	for mainVersion := range result {
		mainVersions = append(mainVersions, mainVersion)
	}
	sort.Strings(mainVersions)
	if len(args) > 0 {
		if _, ok := result[args[0]]; ok {
			fmt.Println("    ", args[0])
			for _, subVersion := range result[args[0]] {
				fmt.Println("\t", subVersion)
			}
		}
	} else {
		for _, mainVersion := range mainVersions {
			fmt.Println("    ", mainVersion)
			for _, subVersion := range result[mainVersion] {
				fmt.Println("\t", subVersion)
			}
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
