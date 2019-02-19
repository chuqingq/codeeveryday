cd mylib
cargo build --release
cd ../mylib2
cargo build --release
cd ..

find . -name "*.so"

export LD_LIBRARY_PATH=`pwd`/mylib/target/release/:`pwd`/mylib2/target/release/
gcc *.c -o test_c_call_rustlib -Lmylib/target/release/ -Lmylib2/target/release/ -lmylib -lmylib2
./test_c_call_rustlib

