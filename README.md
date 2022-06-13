# go-reverse-shell
This is a Go program which will create a reverse shell by connecting to a specified address and port. It was originally created by ILightThings (see below) and was edited by me to add platform-specific functionality as well as a partially stabilized Linux shell.

## Usage
Edit line 16 in shell.go to use the IP address and port of your listener. Compile and then execute the program from the target machine to have it establish a session with you.

Compiling for Windows: `GOOS=windows go build shell.go`

Compiling for Linux: `GOOS=linux go build shell.go`

## Credit
Originally created by:
https://github.com/ILightThings/
