#[cfg(some_condition)]
fn conditional_function() {
        println!("condition met!");
}

fn main() {
        conditional_function();
        if cfg!(some_condition) {
            println!("chuqq match")
        }
}
