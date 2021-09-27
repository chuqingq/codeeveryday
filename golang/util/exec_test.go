package util

import (
	"log"
	"testing"
)

func TestExecCommand(t *testing.T) {
	cmd := "ls -l /dev/disk/by-id 2>/dev/null | grep -v part"
	output, err := ExecCommand("sh", "-c", cmd)
	if err != nil {
		log.Printf("error: %v", err)
		t.Fatalf("exec command error: %v", err)
	}
	log.Printf("output: %v", output)
}

func TestExecShell(t *testing.T) {
	shell := "ls -l /dev/disk/by-id 2>/dev/null | grep -v part"
	output, err := ExecShell(shell)
	if err != nil {
		log.Printf("error: %v", err)
		t.Fatalf("exec shell error: %v", err)
	}
	log.Printf("output: %v", output)
}
