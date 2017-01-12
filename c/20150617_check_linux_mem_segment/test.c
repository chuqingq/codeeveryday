#include <stdio.h>
#include <stdlib.h>

#include <unistd.h>

extern char etext, edata, end; /* The symbols must have some type,
                                   or "gcc -Wall" complains */

int main(int argc, char *argv[])
{
    char *heap;
    char stack;

    printf("First address past:\n");
    printf("    program text (etext)      %p\n", &etext);
    printf("    initialized data (edata)  %p\n", &edata);
    printf("    uninitialized data (end)  %p\n", &end);

    heap = sbrk(0);
    printf("sbrk: %p\n", heap);

    heap = malloc(1);
    printf("malloc: %p\n", heap);

    printf("stack: %p\n", &stack);

    exit(EXIT_SUCCESS);
}
/*
# ./test3
First address past:
    program text (etext)      0x4007d6
    initialized data (edata)  0x601040
    uninitialized data (end)  0x601050
sbrk: 0x1952000
malloc: 0x1952010
stack: 0x7fff6ee0aa0f

结果为：
text
initialized data
uninitialized data
heap(data)
stack
*/
