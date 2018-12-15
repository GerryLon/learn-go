package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// echo
func echo() {
	cmd := exec.Command("echo", "hello world")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", string(output))
}

// ps -ef | grep ssh
func pipeline() {
	ps := exec.Command("ps", "-ef")
	grep := exec.Command("grep", "sshd")
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer

	// ps -> buf1 -> grep -> buf2
	ps.Stdout = &buf1
	grep.Stdin = &buf1
	grep.Stdout = &buf2

	if err := ps.Run(); err != nil {
		fmt.Printf("ps -ef, err:%v\n", err)
		return
	}

	if err := grep.Run(); err != nil {
		fmt.Printf("grep sshd, err:%v\n", err)
		return
	}
	fmt.Printf("ps -ef | grep sshd:\n%s", buf2.String())
}

func main() {
	echo()
	pipeline()
}
