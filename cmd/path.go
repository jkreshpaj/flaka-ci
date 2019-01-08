package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//ServerConfig contains map of services directories
type ServerConfig struct {
	Dir             string
	NotificationURL string
	Services        map[string]map[string]interface{} `yaml:"services"`
}

//Init server config
func (c *ServerConfig) Init(config string, nURL string) error {
	if err := c.SetDir(); err != nil {
		return err
	}
	if err := c.ReadConfig(config); err != nil {
		return err
	}
	if err := c.CheckDirectories(); err != nil {
		return err
	}
	c.NotificationURL = nURL
	return nil
}

//SetDir reads current pwd
func (c *ServerConfig) SetDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	c.Dir = dir
	return nil
}

//ReadConfig reads services from yml file
func (c *ServerConfig) ReadConfig(file string) error {
	data, err := ioutil.ReadFile(c.Dir + "/" + file)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(data), &c); err != nil {
		return err
	}
	return nil
}

//CheckDirectories check if directories exist
func (c *ServerConfig) CheckDirectories() error {
	for _, options := range c.Services {
		path := options["path"].(string)
		if path == "" {
			return errors.New("Error in flaka-ci.yml: path option is required")
		}
		if _, err := os.Stat(c.Dir + "/" + path); os.IsNotExist(err) {
			return errors.New("flaka-ci.yml: Could not find directory " + path + " in " + c.Dir)
		}
	}
	return nil
}
