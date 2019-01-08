package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
)

//UpdateLogRegexp match git pull log
var UpdateLogRegexp = `(?m)Updating [a-z0-9]{7}..[a-z0-9]{7}$`

//Mutex lock/unclock command
var Mutex = &sync.Mutex{}

//PullRepository pulls a service repository
func PullRepository(watcher *Watcher, done chan bool) error {
	checker, err := regexp.Compile(UpdateLogRegexp)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Dir = watcher.ServicePath
	err = cmd.Run()
	if err != nil {
		return err
	}
	updateLog := checker.FindString(string(stdout.Bytes()))
	log.Println("\u001b[33m" + "[i] " + watcher.ServicePath + " " + updateLog + "\u001b[0m")
	if watcher.SendNotification() {
		ntf := Notification{
			EndpointURL: watcher.Notifications,
			Title:       "SUCCESSFUL: Update service *" + watcher.ServiceName + "*",
			Log:         "```Output:\n" + stdout.String() + "```",
			Type:        "success",
		}
		if err := ntf.Send(); err != nil {
			return err
		}
	}
	done <- true
	return nil
}

//ParseCommands return array of string commands from interface
func ParseCommands(commands interface{}) ([]string, error) {
	var cArr []string
	c := reflect.ValueOf(commands)
	for i := 0; i < c.Len(); i++ {
		cArr = append(cArr, c.Index(i).Elem().String())
	}
	return cArr, nil
}

//ExecCommand runs command of the service specified in yml file
func ExecCommand(watcher *Watcher, command string) error {
	Mutex.Lock()
	commands := strings.Split(command, " ")
	var cmd *exec.Cmd
	if len(commands) > 1 {
		cmd = exec.Command(commands[0], commands[1:]...)
	} else {
		cmd = exec.Command(commands[0])
	}
	var stdout, stderr bytes.Buffer
	cmd.Dir = watcher.ServicePath
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	time.Sleep(5 * time.Second)
	err := cmd.Run()
	if err != nil {
		if stderr.String() != "" {
			errstd := stderr.String()
			HandleError("ERROR: Run command: "+strings.Join(commands, " "), errstd)
		}
	}
	if stdout.String() != "" {
		log.Println(stdout.String())
		if watcher.SendNotification() {
			ntf := Notification{
				EndpointURL: watcher.Notifications,
				Title:       "COMPLETED: Run command *" + strings.Join(commands, " ") + "* for service *" + watcher.ServiceName + "*",
				Log:         "```Output:\n" + stdout.String() + "```",
				Type:        "success",
			}
			if err := ntf.Send(); err != nil {
				return err
			}
		}
	}
	Mutex.Unlock()
	return nil
}

//HandleError logs error and sends to slack
func HandleError(text string, errLog string) {
	if server.NotificationURL != "" {
		ntf := Notification{
			EndpointURL: server.NotificationURL,
			Title:       text,
			Log:         "```" + errLog + "```",
			Type:        "error",
		}
		if err := ntf.Send(); err != nil {
			log.Println(err)
		}
	}
	log.Println(text)
	log.Println(errLog)
}
