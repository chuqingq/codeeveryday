#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>

int fd;

int random() {
    int random;
    int ret = read(fd, &random, sizeof(random));
    if (ret != sizeof(random)) {
        printf("ret: %d\n", ret);
        return -1;
    }
    return 0;
}

int main() {
    fd = open("/dev/urandom", O_RDONLY);
    if (fd <= 0) {
        perror("open file error");
        return -1; 
    }

    for (int i = 0; i < 10000; i++) {
        random();
    }
    return 0;
}

