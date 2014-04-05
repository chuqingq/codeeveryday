#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h> // inet_addr
#include <unistd.h>

#include <stdio.h>
#include <string.h>

int main()
{
  int ret;

  int client = socket(AF_INET, SOCK_STREAM, 0);
  if (client < 0) {perror("socket error"); return -1;}

  struct sockaddr_in addr;
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr("127.0.0.1");
  addr.sin_port = htons(8890);

  ret = connect(client, (struct sockaddr*)(&addr), sizeof(struct sockaddr));
  if (ret < 0) {perror("connect error"); return -1;}

  char buf[64];
  ret = read(client, buf, sizeof(buf));
  if (ret < 0) {perror("read error"); return -1;}

  close(client);

  return 0;
}

