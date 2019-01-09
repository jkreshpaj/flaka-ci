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
		log.Print("FlakaCI started in background")

		// if cmd.Flag("detach").Value.String() == true {
		// 	cntxt := &daemon.Context{
		// 		PidFileName: "flaka-pid",
		// 		PidFilePerm: 0644,
		// 		LogFileName: "flaka-log",
		// 		LogFilePerm: 0640,
		// 		WorkDir:     "./",
		// 		Umask:       027,
		// 		Args:        []string{"[go-daemon sample]"},
		// 	}
		// 	d, err := cntxt.Reborn()
		// 	if err != nil {
		// 		log.Fatal("Unable to run: ", err)
		// 	}
		// 	if d != nil {
		// 		return
		// 	}
		// 	defer cntxt.Release()
		// }

		configFile := cmd.Flag("config").Value.String()
		nURL := cmd.Flag("notify").Value.String()
		port := cmd.Flag("port").Value.String()

		if err := server.Init(configFile, nURL); err != nil {
			log.Fatal(err)
		}

		WatchCommits(&server)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", "FlakaCI server running")
		})

		http.ListenAndServe(":"+port, nil)
	},
}

//Execute main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
