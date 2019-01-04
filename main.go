package main

import (
	"github.com/flakaal/flaka-ci/cmd"
)

func main() {

	cmd.Execute()

	// hashReg, _ := regexp.Compile(`(?m)[a-z0-9]{40}\srefs/heads/master$`)

	// cmd := exec.Command("git", "show-ref", "--head")
	// var cmdOut bytes.Buffer
	// cmd.Stdout = &cmdOut
	// cmd.Run()

	// cmd2 := exec.Command("git", "ls-remote")
	// var cmd2Out bytes.Buffer
	// cmd2.Stdout = &cmd2Out
	// cmd2.Run()

	// fmt.Println("[LOCAL REPO COMMIT]:", hashReg.FindAllString(string(cmdOut.Bytes()), 1))
	// fmt.Println("[LOCAL REMOTE COMMIT]:", hashReg.FindAllString(string(cmd2Out.Bytes()), 1))

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)

	// data, err := ioutil.ReadFile(dir + "flaka-ci.yml")

	// s := Services{}

	// if err := yaml.Unmarshal([]byte(data), &s); err != nil {
	// 	log.Fatal(err)
	// }
}
