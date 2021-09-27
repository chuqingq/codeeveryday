package util

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestCapture(t *testing.T) {
	a := "a,b,c"
	r := Capture(a, ",", ",")
	if r != "b" {
		t.Fatalf("Capture() error")
	}
}

func TestCapture2(t *testing.T) {
	a := "a,b,c"
	r := Capture(a, "^", "c")
	log.Printf("r: %v", r)
	if r != "a,b," {
		t.Fatalf("Capture() error")
	}
	r = Capture(a, "b", "$")
	if r != ",c" {
		t.Fatalf("capture 2.2 error")
	}
}

func TestCapture3(t *testing.T) {
	a := "total 0\nlrwxrwxrwx 1 root root 13 Jul 21 10:45 mmc-G8GTF4R_0x74dbd49e -> ./../mmcblk2\n"
	r := Capture(a, "0x", " ")
	if r != "74dbd49e" {
		t.Fatalf("capture result not expected")
	}
	log.Printf("r: %v", r)
	h, err := hex.DecodeString(r)
	log.Printf("h: %v, %v", h, err)
	if err != nil {
		t.Fatalf("hex decode error: %v", err)
	}
}
