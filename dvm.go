package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	log "github.com/Sirupsen/logrus"
)

var version string

func init() {
	// todo: make it configurable through the command line
	log.SetLevel(log.InfoLevel)
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "dvm",
		Long: "dotnet Version Manager (https://github.com/mangostano/dvm).",
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion, _ := cmd.Flags().GetBool("version"); !showVersion {
				return pflag.ErrHelp
			}
			fmt.Println(version)
			return nil
		},
	}

	rootCmd.Flags().Bool("version", false, "version of dvm ")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1);
	}
}