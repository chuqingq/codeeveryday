## function in function

def my_func():
    def my_func1():
        print 'my_func1()'
    class A():
        pass
    pass

if __name__ == '__main__':
    a = A()
    pass