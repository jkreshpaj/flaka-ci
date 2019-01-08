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
	rootCmd.PersistentFlags().StringP("port", "p", "7000", "FlakaCI server port")
	rootCmd.PersistentFlags().StringP("config", "c", "flaka-ci.yml", "Configuration file")
	rootCmd.PersistentFlags().StringP("notify", "n", "", "Webhook url to send automatic log messages")
	rootCmd.PersistentFlags().BoolP("detach", "d", false, "Detached mode runs FlakaCI in background")
}

var rootCmd = &cobra.Command{
	Use:   "flaka-ci",
	Short: "Run flaka-ci [arg] to start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.Init(cmd.Flag("config").Value.String()); err != nil {
			log.Fatal(err)
		}

		fmt.Println(cmd.Flag("detach").Value)

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
