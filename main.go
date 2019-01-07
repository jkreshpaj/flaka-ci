package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Dir      string
	Services map[string]map[string]interface{} `yaml:"services"`
	Port     int                               `yaml:"port"`
}

func main() {
	c := ServerConfig{}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(dir + "/flaka-ci.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal([]byte(data), &c); err != nil {
		log.Fatal(err)
	}

	// arr := reflect.ValueOf(c.Services["test"]["command"])
	// for i := 0; i < arr.Len(); i++ {
	// 	fmt.Printf(arr.Index(i).Elem().String())
	// }

	// cmd.Execute()
}
