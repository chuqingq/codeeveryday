use std::io::Read;

#[derive(Debug)]
struct Foo {
    a: i32,
}

/*
impl Read for Foo {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize> {
        Ok(0)
    }
}
*/

impl Read for &mut Foo {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize> {
        self.a = 10;
        Ok(0)
    }
}

fn bar<T: Read>(mut r: T) {
    let mut buf = [0; 10];
    r.read(&mut buf[..]);
    println!("123");
}

fn main() {
    let mut f = Foo{a: 1};
    //bar(f);
    bar(&mut f); // the trait `std::io::Read` is not implemented for `&Foo`
}
