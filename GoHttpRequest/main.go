package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {
	url := "https://raw.githubusercontent.com/suzakutakumi/MarkdownParser/master/command/MarkdownParser.py"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	cmd := exec.Command("python3", "-c", string(byteArray))

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(stdin, "# aaa\n- a\n-b\n 1. aaa\n 2. bbb")

	stdin.Close()

	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
