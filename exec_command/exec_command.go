package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	output, err := exec.Command("ls", "-la").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	println(string(output))
}
