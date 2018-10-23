import ctypes
from ctypes import *

class Version(Structure):
    _fields_ = [
        ("cCgtVersion", c_char * 64),
        ("cReleaseVer", c_char * 64),
        ("ncReleaseTime", c_char * 64)
    ]

if __name__ == '__main__':
    lib = ctypes.cdll.LoadLibrary("./libTTSEngine.so")
    version = Version()
    lib.iMedia_CTTS_GetVersion(pointer(version))
    print "cCgtVersion:", version.cCgtVersion
    print "cReleaseVer:", version.cReleaseVer
    print "ncReleaseTime:", version.ncReleaseTime

