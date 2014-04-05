package main

import (
    "net"
    "fmt"
    "io"
)

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:8889")//侦听在6666端口
    if err != nil {
        panic("error listening:"+err.Error())
    }
    fmt.Println("Starting the server")
 
    for {
        conn, err := listener.Accept()
        if err != nil {
            panic("Error accept:"+err.Error())
        }
        fmt.Println("Accepted the Connection :", conn.RemoteAddr())
        go EchoServer(conn)
    }
}

func EchoServer(conn net.Conn) {
    buf := make([]byte, 1024)
    defer conn.Close()

    for {
        n, err := conn.Read(buf);
        switch err {
            case nil:
                conn.Write( buf[0:n] )
            case io.EOF:// notice
                fmt.Printf("Warning: End of data: %s \n", err);
                return
            default:
                fmt.Printf("Error: Reading data : %s \n", err);
                return
        }
     }
}
