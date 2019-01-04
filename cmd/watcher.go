package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var CommitHashRegexp = `(?m)[a-z0-9]{40}\srefs/heads/master$`

type Watcher struct {
	ServiceName string
	ServicePath string
}

func (w *Watcher) Start() error {
	for {
		hash, err := w.LocalMasterHash()
		if err != nil {
			return err
		}
		remoteHash, err := w.RemoteMasterHash()
		if err != nil {
			return err
		}
		if hash != remoteHash {
			fmt.Println("REMOTE HAS CHANGED")
		}
		return nil
	}
}

func (w *Watcher) LocalMasterHash() (string, error) {
	checker, err := regexp.Compile(CommitHashRegexp)
	if err != nil {
		return "", err
	}
	cmd := exec.Command("git", "show-ref", "--head")
	cmd.Dir = w.ServicePath
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Run()
	match := checker.FindString(string(cmdOut.Bytes()))
	hash := strings.Split(match, " ")[0]
	return hash, nil
}

func (w *Watcher) RemoteMasterHash() (string, error) {
	checker, err := regexp.Compile(CommitHashRegexp)
	if err != nil {
		return "", err
	}
	cmd := exec.Command("git", "ls-remote")
	cmd.Dir = w.ServicePath
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Run()
	match := checker.FindString(string(cmdOut.Bytes()))
	hash := strings.Split(match, "\t")[0]
	return hash, nil
}

func WatchCommits(c *ServerConfig) {
	for key, value := range c.Services {
		w := new(Watcher)
		w.ServiceName = key
		w.ServicePath = c.Dir + "/" + value
		if err := w.Start(); err != nil {
			log.Fatal(err)
		}
	}
}
