#include <stdio.h>
#include <strings.h>
#include <assert.h>

#include <unistd.h>
#include <fcntl.h> // fcntl

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include <sys/epoll.h>

static int set_nonblock(int fd) {
  int flags, ret;
  flags = fcntl(fd, F_GETFL, 0);
  assert(flags >= 0);
  ret = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
  assert(ret >= 0);
  return 0;
}

int main() {
  int server, ret;

  server = socket(AF_INET, SOCK_STREAM, 0);
  assert(server >= 0);

  struct sockaddr_in addr;
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_addr.s_addr = inet_addr("0.0.0.0");
  addr.sin_port = htons(8890);

  int reuse = 1;
  ret = setsockopt(server, SOL_SOCKET, SO_REUSEADDR, &reuse, sizeof(reuse));
  assert(ret == 0);

  ret = bind(server, (struct sockaddr*)&addr, sizeof(struct sockaddr));
  if (ret < 0) {perror("bind"); return -1;}

  ret = listen(server, 5);
  if (ret < 0) {perror("listen"); return -1;}

  ret = set_nonblock(server);
  assert(ret == 0);

  int epfd = epoll_create(1024);
  if (epfd < 0) {perror("epoll_create"); return -1;}

#define MAX 10
  struct epoll_event ev[MAX];

  ev[0].events = EPOLLIN;
  ev[0].data.u64 = 0;
  ev[0].data.fd = server;
  ret = epoll_ctl(epfd, EPOLL_CTL_ADD, server, &ev[0]);
  assert(ret == 0);

  char buf[64];
  size_t len = 0;
  while(1) {
    ret = epoll_wait(epfd, ev, MAX, -1);
    assert(ret >= 0);

    for (int i = 0; i < ret; ++i) {
      if (ev[i].events & EPOLLIN) {
        if (ev[i].data.fd == server) {
          puts("EPOLLIN server");
          socklen_t addrlen = sizeof(struct sockaddr);
          int s = accept(server, (struct sockaddr*)&addr, &addrlen);
          assert(ret >= 0);
          ret = set_nonblock(s);
          assert(ret == 0);
          ev[i].events = EPOLLIN;
          ev[i].data.u64 = 0;
          ev[i].data.fd = s;
          ret = epoll_ctl(epfd, EPOLL_CTL_ADD, s, &ev[i]);
          assert(ret == 0);
          continue;
        }
        puts("EPOLLIN");
        ret = read(ev[i].data.fd, buf, sizeof(buf));
        assert(ret >= 0);
        if (ret == 0) {
          ret = epoll_ctl(epfd, EPOLL_CTL_DEL, ev[i].data.fd, &ev[i]);
          assert(ret >= 0);
          close(ev[i].data.fd);
        }
        else {
          len = ret;
          ev[i].events = EPOLLIN | EPOLLOUT;
          ret = epoll_ctl(epfd, EPOLL_CTL_MOD, ev[i].data.fd, &ev[i]);
          assert(ret == 0);
        }
      }
      else if (ev[i].events & EPOLLOUT) {
        puts("EPOLLOUT");
        if (len > 0) {
          ret = write(ev[i].data.fd, buf, len);
          assert(ret == len);
          len = 0;
          ev[i].events = EPOLLIN;
          ret = epoll_ctl(epfd, EPOLL_CTL_MOD, ev[i].data.fd, &ev[i]);
          assert(ret >= 0);
        }
      }
      else if (ev[i].events & EPOLLERR) {
        puts("EPOLLERR");
      }
      else if (ev[i].events & EPOLLHUP) {
        puts("EPOLLHUP");
      }
    }
  }

  return 0;
}
