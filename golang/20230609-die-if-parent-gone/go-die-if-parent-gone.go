package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Demonstration of "forking" and setting prctl PDEATH control
// Note: This is not practical since you don't really want to do
// fork/exec in golang. That will really mess things up due to the
// way the thread scheduler works.

// PR_SET_PDEATHSIG is the signal we will use with the prctl syscall so
// the thread (links are made to the parent thread, not necessarily the parent
// process) that forked children will exit when given a SIGKILL signal.
// You don't have to set it to SIGKILL, but I did.

// ####################################################
// #      HOW TO VERIFY CORRECT BEHAVIOR              #
// ####################################################
// In a separate terminal, run 'ps -a' and verify you have 2 'main' processes.
// Then run 'kill <parent pid>' and verify that the child is also killed.

// Next, comment out line 66 of this code and rebuild/rerun and perform the
// same steps. You'll see that the child should survive and needs to be killed manually

// NOTE: Due to runtime behavior, using ^C on the terminal running this process, it will also
// cleanup the processes tied to this parent process. Cheers :)

const (
	PRCTL_SYSCALL    = 157
	PR_SET_PDEATHSIG = 1
)

func fork() {
	cmd := exec.Command(os.Args[0])

	// Set the parent process id as an ENVVAR so we can handle the race condition possibility
	// (which arises due to the delay between "forking" and setting the PDEATH control)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PARENT_PID=%d", os.Getpid()))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	fmt.Println("Started child process", cmd.Process.Pid)
}

func setKillSignal() {
	_, _, errno := syscall.RawSyscall(uintptr(PRCTL_SYSCALL), uintptr(PR_SET_PDEATHSIG), uintptr(syscall.SIGKILL), 0)
	if errno != 0 {
		os.Exit(127 + int(errno))
	}
	// here's the check that prevents an orphan due to the possible race
	// condition
	if strconv.Itoa(os.Getppid()) != os.Getenv("PARENT_PID") {
		os.Exit(1)
	}
}

func DieIfParentDie() {
	_, _, errno := syscall.RawSyscall(uintptr(PRCTL_SYSCALL), uintptr(PR_SET_PDEATHSIG), uintptr(syscall.SIGTERM), 0)
	if errno != 0 {
		log.Printf("prctl PR_SET_PDEATHSIG error: %v", errno)
	}
	log.Printf("DieIfParentDie: child: %v, parent: %v", os.Getpid(), os.Getppid())
}

func main() {
	if os.Getenv("PARENT_PID") != "" {
		// Comment the line below
		// DieIfParentDie()
	} else {
		fmt.Println("Parent ID", os.Getpid())
		fork()
		time.Sleep(time.Second * 2)
		panic("parent panic")
	}
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan)
	sig := <-killChan
	log.Printf("%v recv sig: %v", os.Getpid(), sig)
}

// 1. 如果不设置prctl，父进程退出，子进程不会收到任何信号
// 2. prctl可以设置父进程退出时发送SIGKILL或SIGTERM信号
// 3. 如果parent panic，也遵循上述结论。


