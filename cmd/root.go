package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var server = ServerConfig{}

func init() {
	rootCmd.PersistentFlags().String("port", "7000", "FlakaCI server port")
}

var rootCmd = &cobra.Command{
	Use:   "flaka-ci",
	Short: "Run flaka-ci [arg] to start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.Init(); err != nil {
			log.Fatal(err)
		}

		WatchCommits(&server)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", "FlakaCI server running")
		})

		http.ListenAndServe(":"+cmd.Flag("port").Value.String(), nil)
	},
}

//Execute main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
