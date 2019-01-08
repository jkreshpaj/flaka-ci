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
	ServiceName     string
	ServicePath     string
	ServiceCommands []string
	Notifications   string
}

//Start compares remote and local commit on master
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
			if w.SendNotification() {
				ntf := Notification{
					EndpointURL: w.Notifications,
					Title:       "STARTED: Job started for '" + w.ServiceName + "'",
					Type:        "info",
				}
				if err := ntf.Send(); err != nil {
					return err
				}
			}
			done := make(chan bool)
			go PullRepository(w, done)
			select {
			case <-done:
				if len(w.ServiceCommands) == 0 {
					return nil
				}
				for _, command := range w.ServiceCommands {
					go func(cmd string) {
						if err := ExecCommand(w, cmd); err != nil {
							log.Println(err)
						}
					}(command)
				}
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

//SendNotification has webhook
func (w *Watcher) SendNotification() bool {
	if w.Notifications == "" {
		return false
	}
	return true
}

//WatchCommits creates a new watcher for every service specified
func WatchCommits(c *ServerConfig) {
	for service, options := range c.Services {
		w := new(Watcher)
		w.ServiceName = service
		w.ServicePath = c.Dir + "/" + options["path"].(string)
		w.Notifications = c.NotificationURL
		if options["command"] != nil {
			commands, err := ParseCommands(options["command"])
			if err != nil {
				log.Println(err)
			}
			w.ServiceCommands = commands
		}
		go w.Start()
	}
}
