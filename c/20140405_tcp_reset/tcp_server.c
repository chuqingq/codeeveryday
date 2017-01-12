#include <sys/socket.h>
#include <netinet/in.h> // struct sockaddr_in
#include <arpa/inet.h> // inet_addr
#include <unistd.h>

#include <stdio.h>
#include <string.h>

int main()
{
  int ret;

  struct sockaddr_in addr;
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr("0.0.0.0");
  addr.sin_port = htons(8890);

  int server;
  server = socket(AF_INET, SOCK_STREAM, 0);
  if (server < 0) {perror("socket error"); return -1;}

  int reuse_addr = 1;
  ret = setsockopt(server, SOL_SOCKET, SO_REUSEADDR, &reuse_addr, sizeof(reuse_addr));
  if (ret < 0) {perror("setsockopt error"); return -1;}

  ret = bind(server, (struct sockaddr*)(&addr), sizeof(struct sockaddr));
  if (ret < 0) {perror("bind error"); return -1;}

  ret = listen(server, 5);
  if (ret < 0) {perror("listen error:"); return -1;}

  int s;
  bzero(&addr, sizeof(addr));
  socklen_t len;
  s = accept(server, (struct sockaddr*)(&addr), &len);
  if (s < 0) {perror("accept error"); return -1;}

  struct linger lin;
  lin.l_onoff = 1;
  lin.l_linger = 0;
  ret = setsockopt(s, SOL_SOCKET, SO_LINGER, &lin, sizeof(lin));
  if (ret < 0) {perror("setsockopt error"); return -1;}

  sleep(2);
  // *(int*)0 = 10;
  close(s);

  return 0;
}
