struct Point {
  x: i32,
  y: i32,
}

fn mypoint() -> Point {
  let x = Point{x: 1, y: 2};
  println!("in: 0x{:x}", &x.x as *const i32 as usize);
  println!("in: {:p}", &x.x);
  x
}

fn main() {
  let x = mypoint();
  println!("in: 0x{:X}", &x.x as *const i32 as usize);
  println!("{}, {}", x.x, x.y);
}
