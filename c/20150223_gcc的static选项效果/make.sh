gcc -static -o test_gcc_static test_gcc_static.c
gcc -o test_gcc_nostatic test_gcc_static.c

echo "ldd -r test_gcc_static:" > result.out
ldd -r test_gcc_static >> result.out

echo "ldd -r test_gcc_nostatic;" >> result.out
ldd -r test_gcc_nostatic >> result.out

echo "ls -l test_gcc_{,no}static" >> result.out
ls -l test_gcc_{,no}static >> result.out
