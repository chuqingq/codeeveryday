#include <stdio.h>

int main(int argc, char* argv[]) {
	int num, *p;
	p = &num;
	num = 1;// little: 0x01000000
	
	if(*(unsigned char *)p == 0x01) {
	  printf("little\n");
	}
	else {
	  printf("big\n");
	}
	
	return 0;
}

// 结论：无论是windows还是linux操作系统，x86系列（包括x86_64）cpu都是小端