package main

import "github.com/flakaal/flaka-ci/cmd"

type ServerConfig struct {
	Dir      string
	Services map[string]map[string]interface{} `yaml:"services"`
	Port     int                               `yaml:"port"`
}

func main() {

	// arr := reflect.ValueOf(c.Services["test"]["command"])
	// for i := 0; i < arr.Len(); i++ {
	// 	fmt.Printf(arr.Index(i).Elem().String())
	// }

	cmd.Execute()
}
