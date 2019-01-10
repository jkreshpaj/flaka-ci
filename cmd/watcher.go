package cmd

import (
	"bytes"
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
		if w.HasChanged() {
			w.job()
		}
	}
}

//HasChanged checks if remote is updated
func (w *Watcher) HasChanged() bool {
	hash, err := w.LocalMasterHash()
	if err != nil {
		HandleError("Error getting local hash", err)
	}
	remoteHash, err := w.RemoteMasterHash()
	if err != nil {
		HandleError("Error getting master hash", err)
	}
	if remoteHash != hash {
		return true
	}
	return false
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
				HandleError("", err)
			}
			w.ServiceCommands = commands
		}
		go w.Start()
	}
}

func (w *Watcher) job() error {
	done := make(chan bool)
	w.composeNotification("Started job for service ", "info", "")
	go PullRepository(w, done)
	select {
	case <-done:
		if len(w.ServiceCommands) == 0 {
			return nil
		}
		for _, command := range w.ServiceCommands {
			go func(cmd string) {
				if err := ExecCommand(w, cmd); err != nil {
					HandleError("Error running command", err)
				}
			}(command)
		}
	}
	return nil
}

func (w *Watcher) composeNotification(text string, ntfType string, ntfLog string) {
	if w.SendNotification() {
		ntf := Notification{
			EndpointURL: w.Notifications,
			Title:       text + w.ServiceName,
			Type:        ntfType,
		}
		if ntfLog != "" {
			ntf.Log = ntfLog
		}
		if err := ntf.Send(); err != nil {
			HandleError("Error sending notification", err)
		}
	}
}
