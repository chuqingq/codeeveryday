extern crate async_std;

use async_std::io;
use async_std::net::{TcpListener, TcpStream};
use async_std::prelude::*;
use async_std::task;

async fn process(mut stream: TcpStream) -> io::Result<()> {
    // println!("Accepted from: {}", stream.peer_addr()?);

    let mut length = 0;
    let mut buf = [0; 1024];

    loop {
        // read
        let n = stream.read(&mut buf[length..]).await?;
        length += n;
        if buf[0..length].ends_with(b"\r\n\r\n") {
            // let req = String::from_utf8_lossy(&buf[0..length]).to_string();
            // println!("read: {req}");
            // write
            stream.write_all(b"HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Length: 11\r\n\r\nhello world").await?;
            // println!("write: {resp}");

            length = 0
        }
    }
}

fn main() -> io::Result<()> {
    task::block_on(async {
        let listener = TcpListener::bind("127.0.0.1:8080").await?;
        println!("Listening on {}", listener.local_addr()?);

        let mut incoming = listener.incoming();

        while let Some(stream) = incoming.next().await {
            let stream = stream?;
            task::spawn(async {
                process(stream).await.unwrap();
            });
        }
        Ok(())
    })
}
