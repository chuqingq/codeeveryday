def my_fun(**kw):
    print 'kw=', kw
    print kw['name']
    ## print kw['test'] ## exception

def my_fun2(name, value, **kw):
    print 'name=', name
    print 'value=', value
    print 'kw=', kw


if __name__ == '__main__':
    my_fun(name=1, value='abc')
    my_fun2(value='abc', name=1,test='helloworld')
