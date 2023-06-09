package util

import (
	"log"
	"os"
	"syscall"
)

const (
	PRCTL_SYSCALL    = 157
	PR_SET_PDEATHSIG = 1
)

func DieIfParentGone() {
	_, _, errno := syscall.RawSyscall(uintptr(PRCTL_SYSCALL), uintptr(PR_SET_PDEATHSIG), uintptr(syscall.SIGKILL), 0)
	if errno != 0 {
		log.Printf("prctl PR_SET_PDEATHSIG error: %v", errno)
	}
	log.Printf("DieIfParentDie: child: %v, parent: %v", os.Getpid(), os.Getppid())
}
