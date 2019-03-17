package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"sort"
)

const versionsUrl = "https://raw.githubusercontent.com/mangostano/dvm/develop/config/versions.json"

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "This command to get all of the dotnet core sdk versions",
	Long: `examples of using this command. For example:
		dvm listAll`,
	Run: func(cmd *cobra.Command, args []string) {
		result := make(map[string][]string)
		getVersionJsonFile(versionsUrl, result)
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

func getVersionJsonFile(url string, result map[string][]string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected http GET status: ", resp.Status)
	}

	decodeError := json.NewDecoder(resp.Body).Decode(&result)
	if decodeError != nil {
		log.Fatal("cannot decode JSON: ", err)
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
