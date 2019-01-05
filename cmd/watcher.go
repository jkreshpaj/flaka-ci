package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//CommitHashRegexp match master commit hash
var CommitHashRegexp = `(?m)[a-z0-9]{40}\srefs/heads/master$`

//Watcher type
type Watcher struct {
	ServiceName string
	ServicePath string
}

//Start compares master and local commit on master
func (w *Watcher) Start() error {
	for {
		log.Println("watching", w.ServicePath)
		hash, err := w.LocalMasterHash()
		if err != nil {
			return err
		}
		remoteHash, err := w.RemoteMasterHash()
		if err != nil {
			return err
		}
		if hash != remoteHash {
			err := PullRepository(w.ServicePath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

//LocalMasterHash returns current local commit hash
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

//RemoteMasterHash returns latest remote master hash
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

//WatchCommits creates a new watcher for every service specified
func WatchCommits(c *ServerConfig) {
	for key, value := range c.Services {
		w := new(Watcher)
		w.ServiceName = key
		w.ServicePath = c.Dir + "/" + value
		go w.Start()
	}
}
