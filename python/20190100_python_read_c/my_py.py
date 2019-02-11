
import ctypes
mylib = ctypes.cdll.LoadLibrary("./my_c.so")
mylib.my_c.restype=ctypes.c_char_p
mylib.my_c()
# >>> b'1234567890'
