package main

import (
	"github.com/flakaal/flaka-ci/cmd"
	"net/http"
	"fmt"
)

func main() {

	cmd.Execute()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "%s", "FlakaCI server running")
	})

	http.ListenAndServe(":9999", nil)

}
