package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// 1. exec.LookPath()
	cmdPath, _ := exec.LookPath("ls")
	fmt.Printf("cmdPath=%s\n", cmdPath)

	// 2. exec.Command()
	fmt.Println()
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	// Run 和 Start只能用一个
	// Run starts the specified command and waits for it to complete.
	_ = cmd.Run()

	// Start starts the specified command but does not wait for it to complete.
	// _ = cmd.Start()
	// _ = cmd.Wait()

	// 3. pipe
	fmt.Println()
	ps := exec.Command("ps", "-ef")
	grep := exec.Command("grep", "-i", "ssh")

	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()
	// ps | grep
	ps.Stdout = w
	grep.Stdin = r

	var buffer bytes.Buffer
	grep.Stdout = &buffer

	_ = ps.Start()
	_ = grep.Start()
	ps.Wait()
	w.Close()
	grep.Wait()
	io.Copy(os.Stdout, &buffer)
}
