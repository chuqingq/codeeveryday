import importlib

def my_import(name):
    mod = importlib.import_module(name)
    mod.my_func()

if __name__ == '__main__':
    my_import('my_module')
