// gcc -o test_reuseaddr{,.c}
#include <netinet/in.h>   
#include <sys/socket.h>   
#include <time.h>   
#include <stdio.h>   
#include <string.h>   
#include <stdlib.h>  
#include <unistd.h>  

int main() {
    int fd, ret, i;
    int reuse = 1;

    struct sockaddr_in destaddr;
    destaddr.sin_family = AF_INET;
    destaddr.sin_port = htons(18080);
    ret = inet_pton(AF_INET, "192.168.23.72", &destaddr.sin_addr);
    if (ret != 1) {
        perror("inet_pton() error");
        return -1;
    }

    struct sockaddr_in localaddr;
    bzero(&localaddr, sizeof(localaddr));
    localaddr.sin_family = AF_INET;
    localaddr.sin_port = htons(20020);
    localaddr.sin_addr.s_addr = htonl(INADDR_ANY);
    /*ret = inet_pton(AF_INET, "192.168.23.72", &localaddr.sin_addr);
    if (ret != 1) {
        perror("local inet_pton()  error");
        return -1;
    }*/

    for (i = 0; i < 65536; i++) {
        fd = socket(AF_INET, SOCK_STREAM, 0);
        if (fd == -1) {
            perror("socket() error");
            return -1;
        }

        ret = setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &reuse, sizeof(reuse));
        if (ret != 0) {
            perror("setsockopt() error");
            return -1;
        }

        ret = bind(fd, (struct sockaddr *)&localaddr, sizeof(localaddr));
        if (ret != 0) {
            perror("bind() error");
            return -1;
        }

        ret = connect(fd, (struct sockaddr *)&destaddr, sizeof(destaddr));
        if (ret != 0) {
            perror("connect() error");
            return -1;
        }

        close(fd);
        printf("count=%d\n", i);
    }

}
