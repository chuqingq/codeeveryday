#include <stdio.h>

#include "test1.h"

int main() {
    my_set(1, 2);
    printf("%d\n", my_get(1));
    return 0;
}

/*
现在新版本gcc的机器上打包libmylib.a，然后编译出依赖mylib的test1.o（c++实现，接口是C的）；
把test1.h和test1.o拷贝到旧版本gcc的机器上，和main.c一起链接。
结果失败：
[root@localhost ~]# gcc main.c test1.o -I. -L. -lmylib
/usr/bin/ld: ./libmylib.a(functexcept.o): unrecognized relocation (0x2a) in section `.text._ZSt21__throw_bad_exceptionv'
/usr/bin/ld: final link failed: 错误的值
collect2: 错误：ld 返回 1
初步分析，可能是bintuils中的ld版本不一致。
*/
