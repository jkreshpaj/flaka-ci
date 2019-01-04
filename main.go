package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("git", "show-ref --head")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	fmt.Println(string(out.Bytes()))
}
