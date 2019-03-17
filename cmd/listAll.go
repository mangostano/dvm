package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "This command to get all of the dotnet core sdk versions",
	Long: `examples of using this command. For example:
		dvm list-all`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, err := os.Open("/Users/lqqu/go/src/dotnet-manager/dvm/config/versions.json")
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		result := make(map[string][]string)
		json.Unmarshal([]byte(byteValue), &result)

		var mainVersions []string
		for mainVersion := range result {
			mainVersions = append(mainVersions, mainVersion)
		}
		sort.Strings(mainVersions)
		for _, mainVersion := range mainVersions {
			fmt.Println("    ", mainVersion)
			for index := range result[mainVersion] {
				fmt.Println("\t", result[mainVersion][index])
			}
		}
	},
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
