#ifndef CLIBRARY_H
#define CLIBRARY_H
//定义函数指针
typedef int (*callback_fcn)(int);
void some_c_func(callback_fcn);
#endif
