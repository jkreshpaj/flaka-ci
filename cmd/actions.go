package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
)

//UpdateLogRegexp match git pull log
var UpdateLogRegexp = `(?m)Updating [a-z0-9]{7}..[a-z0-9]{7}$`

//PullRepository pulls a service repository
func PullRepository(path string, done chan bool) error {
	checker, err := regexp.Compile(UpdateLogRegexp)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return err
	}
	updateLog := checker.FindString(string(stdout.Bytes()))
	log.Println("\u001b[33m" + "[i] " + path + " " + updateLog + "\u001b[0m")
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
func ExecCommand(path string, command string, done chan bool) error {
	commands := strings.Split(command, " ")
	var cmd *exec.Cmd
	if len(commands) > 1 {
		cmd = exec.Command(commands[0], commands[1:]...)
	} else {
		cmd = exec.Command(commands[0])
	}
	var stdout, stderr bytes.Buffer
	cmd.Dir = path
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("An error has occured while running command", command)
		if stderr.String() != "" {
			errstd := strings.Split(stderr.String(), "\n")[0]
			log.Println(errstd)
		}
	}
	if stdout.String() != "" {
		log.Println(stdout.String())
	}
	done <- true
	return nil
}
