package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

//UpdateLogRegexp match git pull log
var UpdateLogRegexp = `(?m)Updating [a-z0-9]{7}..[a-z0-9]{7}$`

//PullRepository pulls a service repository
func PullRepository(path string) error {
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
	log.Println("\u001b[33m" + "[i] " + updateLog + "\u001b[0m")
	return nil
}
