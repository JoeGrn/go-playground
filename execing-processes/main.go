package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// this replaces go process with ls process
	// rather than spawning a new process
	// it requires an absolute path in this case to bin/ls
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// args for ls command in slice format
	args := []string{"ls", "-a", "-l", "-h"}

	// also needs env variables
	env := os.Environ()

	// the actual syscall
	// if it is sucessful execution of the go process ends here and it is replaces with /bin/ls -a -l -h
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	fmt.Println("This will not be printed")
}
