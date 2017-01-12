#include <stdio.h>
#include <string.h>

#include <unistd.h>

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

int main()
{
  int client, ret;
  client = socket(AF_INET, SOCK_STREAM, 0);
  if (client < 0) {perror("socket error"); return -1;}

  struct sockaddr_in addr;
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr("127.0.0.1");
  addr.sin_port = htons(8890);

  ret = connect(client, (struct sockaddr*)&addr, sizeof(struct sockaddr));
  if (ret < 0) {perror("connect error"); return -1;}

  const char* msg = "hello world";
  size_t len = strlen(msg);
  ret = write(client, msg, len);
  if (ret != len) {printf("write %d\n", ret); return -1;}

  char buf[64];
  ret = read(client, buf, sizeof(buf));
  if (ret < 0) {perror("read error"); return -1;}
  buf[ret] = '\0';
  printf("read %d: %s\n", ret, buf);

  close(client);
  return 0;
}

