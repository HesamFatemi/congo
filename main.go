package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		log.Panic("command not found.")
	}
}

func run() {
	fmt.Printf("debug run func: runnig %v, with PID %v", os.Args[2:], os.Getpid())

	// passing the command to os
	c := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	c.Run()
}

func child() {
	fmt.Printf("debug child func: runnig %v, with PID %v", os.Args[2:], os.Getpid())

	//set hostname for the container

	err := syscall.Sethostname([]byte("congo"))
	if err != nil {
		fmt.Printf("faild to cahnge hostname err: %v", err)
	}

	// passing the command to oss
	c := exec.Command(os.Args[2], os.Args[3:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	c.Run()
}
