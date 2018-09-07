fn make_vec() -> Vec<i32> {
    let mut vec = Vec::new();
    vec.push(0);
    vec.push(1);
    vec // transfer ownership to the caller
}


fn print_vec(vec: &Vec<i32>) {
    // the `vec` parameter is borrowed for this scope
 
    for i in vec.iter() {
        println!("{}", i)
    }
 
    // now, the borrow ends
}

fn push_all(from: &Vec<i32>, to: &mut Vec<i32>) {
    for i in from.iter() {
        to.push(*i);
    }
}
 
fn use_vec() {
    let mut vec = make_vec();
    push_all(&vec, &mut vec);
    // for i in vec.iter() {  // continue using `vec`
    //     println!("{}", i * 2)
    // }
    println!("{:?}", vec)
    // vec is destroyed here
}

fn main() {
    use_vec()
}

/*
error[E0502]: cannot borrow `vec` as mutable because it is also borrowed as immutable
  --> test_borrow.rs:27:25
   |
27 |     push_all(&vec, &mut vec);
   |               ---       ^^^- immutable borrow ends here
   |               |         |
   |               |         mutable borrow occurs here
   |               immutable borrow occurs here

error: aborting due to previous error
*/
