#[no_mangle]
pub extern "C" fn foo(a: i32, b: i32) {
    println!("hello : a + b = {}", a + b);
}

