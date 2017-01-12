#define MINGW32
 
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
 
#ifdef MINGW32
#include <winsock2.h>
#else
#include <sys/socket.h>
#include <arpa/inet.h>
#endif
 
#define LISTENQ 10
 
int main(int argc,char *argv[])
{
  int serverfd,connectfd;
  struct sockaddr_in serveraddr;
  char buff[1024];
  time_t tlick;
  int iRet;
 
  #ifdef MINGW32
  //Winsows下启用socket
  WSADATA wsadata;
  if(WSAStartup(MAKEWORD(1,1),&wsadata)==SOCKET_ERROR)
  {
    printf("WSAStartup() fail\n");
    exit(0);
  }
  #endif
 
  //新建socket
  printf("socket()\n");
  serverfd=socket(AF_INET,SOCK_STREAM,0);
  if(serverfd==-1)
  {
    printf("socket() fail\n");
    exit(0);
  }
 
  //清零
  memset(&serveraddr,0,sizeof(serveraddr));
 
  //设置协议
  serveraddr.sin_family=AF_INET;
  //设置IP
  serveraddr.sin_addr.s_addr=htonl(INADDR_ANY);
  //设置Port
  serveraddr.sin_port=htons(1024);
 
  //绑定端口，监听1024端口的任何请求
  printf("bind()\n");
  iRet=bind(serverfd,(struct sockaddr*)&serveraddr,sizeof(serveraddr));
  if(iRet==-1)
  {
    printf("bind() fail\n");
    exit(0);
  }
 
  //监听端口，最大并发数10
  printf("listen()\n");
  iRet=listen(serverfd,LISTENQ);
  if(iRet==-1)
  {
    printf("listen() fail\n");
    exit(0);
  }
 
  //接受请求，发送主机时间
  for(;;)
  {
    printf("Waiting for connection...\n");
    //接受请求
    connectfd=accept(serverfd,(struct sockaddr*)NULL,NULL);
    //获取时间 
    tlick=time(NULL);
    //格式化时间 
    snprintf(buff,sizeof(buff),"From mys:\n%s",ctime(&tlick));
    //写入时间
    //关闭请求
    #ifdef MINGW32
    send(connectfd,buff,strlen(buff),0);
    closesocket(connectfd);
    #else
    write(connectfd,buff,strlen(buff));
    close(connectfd);
    #endif
  }
 
  #ifdef MINGW32
  //Winsows下关闭socket
  closesocket(serverfd);
  WSACleanup();
  #endif
 
  //退出
  exit(0);
}