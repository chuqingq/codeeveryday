use std::sync::Mutex;
use std::thread;

fn main() {
    let counter = Mutex::new(0);

    let mut handles = vec![];

    for _ in 0..10 {
        let handle = thread::spawn( || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("result: {}", *counter.lock().unwrap());
}

/*
error[E0373]: closure may outlive the current function, but it borrows `counter`, which is owned by the current function
  --> /data/162789560328357267.rust:10:37
   |
10 |         let handle = thread::spawn( || {
   |                                     ^^ may outlive borrowed value `counter`
11 |             let mut num = counter.lock().unwrap();
   |                           ------- `counter` is borrowed here
help: to force the closure to take ownership of `counter` (and any other referenced variables), use the `move` keyword
   |
10 |         let handle = thread::spawn( move || {
   |                                     ^^^^^^^

error: aborting due to previous error

For more information about this error, try `rustc --explain E0373`.
*/

