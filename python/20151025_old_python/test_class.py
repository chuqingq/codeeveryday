class MyTest:
    """docstring for MyTest"""
    def __init__(self):## 构造函数
        print 'enter __init__'

    def load(self):
        print 'enter _load'
        pass
    
def test(handle, arg1=[], arg2={}):## 可选参数，可以直接指定arg2
    print handle
    handle()## 闭包
    print 'arg1=',arg1
    print 'arg2=', arg2

class _MyCaller:
    """docstring for MyCaller"""
    def __init__(self):
        self.var = 'hello world'
        print 'enter MyCaller.__init__'

    def __call__(self):
        print self.var
        return self.var
        

if __name__ == '__main__':
    a = MyTest()
    t = {'a':1, 'b':2}
    test(a.load, arg2=t)## 直接传入a的load方法，闭包
    print __file__## 模块对应的文件。相对路径

    my_caller = _MyCaller()## 加前缀_无法表示不导出，需要在__all__中说明要导出的内容
    print my_caller()
    