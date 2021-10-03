use std::io::Read;

#[derive(Debug)]
struct Foo;

impl Read for Foo {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize> {
        Ok(0)
    }
}

/*
impl Read for &Foo {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize> {
        Ok(0)
    }
}
*/

fn bar<T: Read>(r: T) {}

fn main() {
    let f = Foo;
    bar(f);
    // bar(&f); // the trait `std::io::Read` is not implemented for `&Foo`
    println!("{:?}", f);
}
