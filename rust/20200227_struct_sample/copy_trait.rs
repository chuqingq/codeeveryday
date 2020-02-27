#[derive(Debug, Copy, Clone)]
struct Book {
}

fn main() {
    let i = Book {};
    let a = i;

    println!("i: {:?}", i);
    println!("a: {:?}", a);
}

