my_var = 'hello world'

def my_fun():
    global my_var
    my_var = '123'

if __name__ == '__main__':
    print my_var
    my_fun()
    print my_var
