#include <stdio.h>
#include <string.h>

#include <unistd.h>

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

int main()
{
  int server, s, ret;
  server = socket(AF_INET, SOCK_STREAM, 0);
  if (server < 0) {perror("socket error"); return -1;}

  struct sockaddr_in addr;
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr("0.0.0.0");
  addr.sin_port = htons(8890);

  ret = bind(server, (struct sockaddr*)&addr, sizeof(struct sockaddr));
  if (ret < 0) {perror("bind error"); return -1;}

  ret = listen(server, 5);
  if (ret < 0) {perror("listen error"); return -1;}

  socklen_t caddrlen;
  s = accept(server, (struct sockaddr*)&addr, &caddrlen);
  if (s < 0) {perror("accept error"); return -1;}

  char buf[64];
  int len;
  while (1) {
    ret = read(s, buf, sizeof(buf));
    if (ret < 0) {perror("recv error"); return -1;}
    if (ret == 0) {close(s); return 0;}
    buf[ret] = '\0';
    printf("recv %s\n", buf);

    len = ret;
    ret = write(s, buf, len);
    if (ret < 0) {perror("write error"); return -1;}
    if (ret != len) {printf("write %d, != %d\n", ret, len); return -1;}
  }

  return 0;
}

