struct HasDrop;

impl Drop for HasDrop {
    fn drop(&mut self) {
        println!("dropping");
    }
}

impl HasDrop {
    fn dodo(&self) {
        println!("do");
    }
}

fn main() {
    HasDrop{}.dodo();
    println!("end");
}
