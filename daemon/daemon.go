package daemon

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
)

//Process FlakaCI background process options
type Process struct {
	Config  string
	Notify  string
	Port    string
	Pid     int
	Logfile string
}

//Start New FlakaCI process
func (p *Process) Start() {
	log.Println("FlakaCI running in background")
	flags := p.mapFlags()
	p.exec(flags)
}

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
