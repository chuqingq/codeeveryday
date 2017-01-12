package main

import (
    "net"
    "strconv"
)

const LocalIP = "0.0.0.0"
const LocalPort = 8888
const RemoteIP = "127.0.0.1"
const RemotePort = 8889

func main() {
    listener, err := net.Listen("tcp", LocalIP + ":" + strconv.Itoa(LocalPort))
    if err != nil {
        panic("error: Listen " + err.Error())
    }

    for {
        lconn, err := listener.Accept()
        if err != nil {
            panic("error: Accept " + err.Error())
        }

        go func() {
            rconn, err := net.Dial("tcp", RemoteIP + ":" + strconv.Itoa(RemotePort))
            if err != nil {
                panic("error: Dial " + err.Error())
            }

            go my_pipe(lconn, rconn)
            go my_pipe(rconn, lconn)
        }()
    }
}

func my_pipe(lconn net.Conn, rconn net.Conn) {
    buf := make([]byte, 512)
    for {
        len, err := lconn.Read(buf)
        if err != nil {
            // err.Error() == "xxxx"// TODO
            println("Read " + err.Error())
            rconn.Close()
            return
        }
        rconn.Write(buf[:len])
    }
}
