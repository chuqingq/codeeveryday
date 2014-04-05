package main

import (
    "net"
    "fmt"
    // "io"
)

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:8888")
    if err != nil {
        panic("error listening:"+err.Error())
    }
    fmt.Println("Starting the server")
 
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic("Error accept:"+err.Error())
        }
        conn.Close()
    }
}
