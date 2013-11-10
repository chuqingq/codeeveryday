package main

import (
    "net"
    // "bufio"
    "fmt"
)

func main() {
    conn, err := net.Dial("tcp", "www.baidu.com:80")
    if err != nil {
        fmt.Printf("dial error: %#v\n", err)
        return
    }

    fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

    data := make([]byte, 1024)
    n, err := conn.Read(data)
    if err != nil {
        fmt.Printf("read error: %#v\n", err)
        return
    }

    fmt.Printf("read: %s\n", data[:n])

    // another way
    // status, err := bufio.NewReader(conn).ReadString('\n')
    // if err != nil {
    //     fmt.Printf("read error: %#v\n", err)
    //     return
    // }

    // fmt.Printf("status: %#v\n", status)
}