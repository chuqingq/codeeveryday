package main
 
import (
    "fmt"
    "time"
    "net"
)

func main() {
    conn,err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()
 
    buf := make([]byte, 1024)
 
    for i := 0; i < 5; i++ {
        //准备要发送的字符串
        msg := fmt.Sprintf("Hello World, %03d", i)
        n, err := conn.Write([]byte(msg))
        if err != nil {
            println("Write Buffer Error:", err.Error())
            break
        }
        fmt.Printf("write: %s\n", msg)
 
        //从服务器端收字符串
        n, err = conn.Read(buf)
        if err !=nil {
            println("Read Buffer Error:", err.Error())
            break
        }
        fmt.Printf("read: %s\n", string(buf[0:n]))
 
        //等一秒钟
        time.Sleep(time.Second)
    }
}
