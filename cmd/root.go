package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var server = ServerConfig{}

var rootCmd = &cobra.Command{
	Use:   "flaka-ci",
	Short: "Run flaka-ci [arg] to start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.Init(); err != nil {
			log.Fatal(err)
		}
		WatchCommits(&server)
	},
}

//Execute main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
