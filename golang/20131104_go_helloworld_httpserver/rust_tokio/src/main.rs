extern crate tokio;

use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpListener;
// use std::env;
use std::error::Error;

const RESPONSE: &[u8] =
    b"HTTP/1.1 200 OK\r\nConnection: keep-alive\r\nContent-Length: 11\r\n\r\nhello world";

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let listener = TcpListener::bind("127.0.0.1:8080").await?;
    println!("Listening on: {}", listener.local_addr()?);

    // accept
    loop {
        let (mut socket, _) = listener.accept().await.expect("accept error");

        tokio::spawn(async move {
            let mut buf = vec![0; 1024];
            let mut length: usize = 0;

            loop {
                // read
                let n = socket
                    .read(&mut buf[length..])
                    .await
                    .expect("failed to read data from socket");

                if n == 0 {
                    return;
                }

                length += n;

                if buf[..length].ends_with(b"\r\n\r\n") {
                    // write
                    let res = socket.write_all(RESPONSE).await;
                    match res {
                        Ok(_) => {
                            // 发送成功
                            length = 0
                        }
                        Err(_) => {
                            println!("write error")
                        }
                    }
                }

                // if let Err(e) = socket.write_all(&buf[..n]).await {
                //     eprintln!("failed to write to socket: {}", e);
                //     return;
                // }
            }
        });
    }
}
