package main

import "testing"

// func main() {
func BenchmarkPingPong(b *testing.B) {
    PingPong(b.N)
}

func PingPong(attempts int) {
    pingCh := make(chan *int)
    pongCh := make(chan *int)
    value := new(int)
    // ponger等待接收
    go ponger(pingCh, pongCh, attempts)
    pinger(value, pingCh, pongCh, attempts)
}

func pinger(value *int, pingCh chan *int, pongCh chan *int, attempts int) {
    for i := 0; i < attempts; i++ {
        *value = i
        pongCh <- value
        msg := <-pingCh
        if *msg != i {
            panic("ping msg invalid")
        }
    }
}

func ponger(pingCh chan *int, pongCh chan *int, attempts int) {
    for i := 0; i < attempts; i++ {
        msg := <-pongCh
        if *msg != i {
            panic("msg invalid")
        }
        pingCh <- msg
    }
}

/*
chuqq@arch-vb~/t/20220803-pingpong-bench $ go test -bench .
goos: linux
goarch: amd64
pkg: pingpong_bench
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkPingPong-6      4335214               295.3 ns/op
PASS
ok      pingpong_bench  1.564s
*/