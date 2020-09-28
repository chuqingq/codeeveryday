package main

/*
#include <stdio.h>
int callOnMeGo(int);

// The gateway function
int callOnMeGo_cgo(int in)
{
  printf("C.callOnMeGo_cgo(): called with arg = %d\n", in);
  //调用GO函数
  return callOnMeGo(in);
}
*/
import "C"
