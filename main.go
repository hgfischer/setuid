package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	runtime.LockOSThread()
	_, _, errNo := syscall.Syscall(syscall.SYS_SETUID, 0, 0, 0)
	if errNo != 0 {
		log.Println(errNo)
	}

	out, err := exec.Command("/usr/bin/id").Output()
	fmt.Println(string(out), err)

	fmt.Println("EUID: ", os.Geteuid())
	fmt.Println("UID: ", os.Getuid())
}
