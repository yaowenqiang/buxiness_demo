package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCommand := exec.Command("grep ", "hello")

	grepIn, _ := grepCommand.StdinPipe()

	grepOut, _ := grepCommand.StdoutPipe()

	grepCommand.Start()
	grepIn.Write([]byte("hello world, are you ok"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)

	grepCommand.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCommand := exec.Command("bash", "-c", "ls -a -l -h")

	lsOut, err := lsCommand.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -l -a -h ")

	fmt.Println(string(lsOut))

}
