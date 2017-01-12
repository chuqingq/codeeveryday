package main
 
import (
    "fmt"
    // "time"
    "net"
)

func main() {
    for i := 1; i < 65536; i++ {
        conn,err := net.Dial("tcp", "127.0.0.1:8888")
        if err != nil {
            panic(err.Error())
        }
        fmt.Printf("%d\n", i)
        conn.Close()
    }
}
