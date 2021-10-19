use std::io::Read;

#[derive(Debug)]
struct Foo {
    a: i32,
}

impl Foo {
    fn write(&mut self) {
        self.a = 20;
    }
}

impl Read for &Foo {
    fn read(&mut self, _buf: &mut [u8]) -> std::io::Result<usize> {
        self.a = 10;
        Ok(0)
    }
}

/*
fn bar<T: Read>(mut r: T) {
    let mut buf = [0; 10];
    r.read(&mut buf[..]);
    println!("123");
}
*/

fn main() {
    let f = Foo{a: 1};
    // bar(&f);
}
