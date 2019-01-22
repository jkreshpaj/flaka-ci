package daemon

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
)

//Process FlakaCI background process options
type Process struct {
	Config      string
	Notify      string
	Port        string
	Pid         int
	UserHomedir string
}

//Start New FlakaCI process
func (p *Process) Start() {
	log.Println("FlakaCI running in background")
	flags := p.mapFlags()
	p.exec(flags)
}

//Kill FlakaCI background process
func (p *Process) Kill() {
	log.Println("KILLING PROCESS", p.Pid)
	processID := strconv.Itoa(p.Pid)
	cmd := exec.Command("kill", processID)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	log.Fatal("KILLED PROCESS", processID)
}

func (p *Process) mapFlags() string {
	var flags string
	flagMap := make(map[string]string)
	flagMap["--config"] = p.Config
	flagMap["--notify"] = p.Notify
	flagMap["--port"] = p.Port
	flags = flags + "flaka-ci "
	for key, value := range flagMap {
		flags = flags + " " + key + " " + value
	}
	return flags
}

func (p *Process) exec(flags string) {
	cmd := exec.Command("screen", "-dmS", "flaka-ci", "bash", "-c", flags)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	if stderr.String() != "" {
		log.Println("ERROR EXEC")
		log.Fatal(stderr.String())
	}
}

//Getpid of FlakaCI process
func (p *Process) Getpid() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	p.UserHomedir = usr.HomeDir
	cmd := exec.Command("/bin/sh", "daemon/pid.sh")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		log.Println("ERROR OCCURED ON RUN")
		log.Println(err)
	}
	if stderr.String() != "" {
		log.Println("ERROR GETPID")
		log.Fatal(stderr.String())
	}
	log.Println(stdout.String())
}

func (p *Process) createDir() {
	if _, err := os.Stat(p.UserHomedir + "/.flaka-ci"); os.IsNotExist(err) {
		if err := os.Mkdir(p.UserHomedir+"/.flaka-ci", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
