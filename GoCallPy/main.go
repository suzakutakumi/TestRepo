package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "big.py")
	currentDir, _ := os.Getwd()
	cmd.Dir = currentDir
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(stdin, "test")
	stdin.Close()

	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
