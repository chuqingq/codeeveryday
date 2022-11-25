use std::thread;
use std::time::Instant;
use std::sync::mpsc::channel;

use crossbeam_channel::bounded;


fn main() {
    stdchannel();
    crossbeamchannel();
}

fn stdchannel() {
    let (s1, r1) = channel();
    let (s2, r2) = channel();

    const LOOPS: i32 = 100000;

    let t = Instant::now();

    // Spawn a thread that receives a message and then sends one.
    thread::spawn(move || {
        for i in 1..LOOPS {
            r1.recv().unwrap();
            s2.send(i).unwrap();
        }
    });

    // Send a message and then receive one.
    for i in 1..LOOPS {
        s1.send(i).unwrap();
        r2.recv().unwrap();
    }

    println!("std-channel: time cost: {:?} ns/loop", t.elapsed().as_nanos()/(LOOPS as u128));
}

fn crossbeamchannel() {
    let (s1, r1) = bounded(0);
    let (s2, r2) = bounded(0);
    const LOOPS: i32 = 10000000;

    let t = Instant::now();

    // Spawn a thread that receives a message and then sends one.
    thread::spawn(move || {
        for i in 1..LOOPS {
            r1.recv().unwrap();
            s2.send(i).unwrap();
        }
    });

    // Send a message and then receive one.
    for i in 1..LOOPS {
        s1.send(i).unwrap();
        r2.recv().unwrap();
    }

    println!("crossbeam-channel: time cost: {:?} ns/loop", t.elapsed().as_nanos()/(LOOPS as u128));
}
// cargo build --release
// ./target/release/pingpong
// std-channel: time cost: 26706 ns/loop // cpu 83%
// crossbeam-channel: time cost: 501 ns/loop // cpu 180%
