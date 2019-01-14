package vendor

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flakaal/flaka-ci/daemon"
	"github.com/spf13/cobra"
)

var server = ServerConfig{}

func init() {
	rootCmd.PersistentFlags().StringP("port", "p", "7000", "FlakaCI server port")
	rootCmd.PersistentFlags().StringP("config", "c", "flaka-ci.yml", "Configuration file")
	rootCmd.PersistentFlags().StringP("notify", "n", "", "Webhook url to send automatic log messages")
	rootCmd.PersistentFlags().BoolP("detach", "d", false, "Detached mode runs FlakaCI in background")
	rootCmd.PersistentFlags().BoolP("stop", "s", false, "Stop daemon process")
}

var rootCmd = &cobra.Command{
	Use:   "flaka-ci",
	Short: "Minimalistic - Zero Configuration CI/CD tool",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := cmd.Flag("config").Value.String()
		nURL := cmd.Flag("notify").Value.String()
		port := cmd.Flag("port").Value.String()

		daemon := daemon.Process{
			Config: configFile,
			Notify: nURL,
			Port:   port,
		}

		if cmd.Flag("detach").Value.String() == "true" {
			daemon.Start()
			os.Exit(0)
		}

		if cmd.Flag("stop").Value.String() == "true" {
			daemon.Kill()
			os.Exit(0)
		}

		if err := server.Init(configFile, nURL); err != nil {
			log.Fatal(err)
		}

		daemon.Pid = os.Getpid()

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
