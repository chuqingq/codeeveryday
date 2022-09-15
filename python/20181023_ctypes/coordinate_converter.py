from ctypes import *

dll = CDLL('../c_lib/libcoorconv.so')
# 设置参数和返回值
dll.coordinate_converter_new.restype = c_void_p

dll.coordinate_convert.argtypes = [c_void_p, c_double, c_double, c_double, POINTER(c_double), POINTER(c_double), POINTER(c_double)]
dll.coordinate_convert.restype = c_void_p

dll.coordinate_converter_free.argtypes = [c_void_p]

# run
converter = dll.coordinate_converter_new(c_double(0), c_double(0), c_double(0))
print(f'converter: {converter}')

lon0 = c_double(5.0)
lat0 = c_double(0)
alt0 = c_double(0)

lon = c_double()
lat = c_double()
alt = c_double()

dll.coordinate_convert(converter, lon0, lat0, alt0, byref(lon), byref(lat), byref(alt))

dll.coordinate_converter_free(converter)

print(f'({lon0.value}, {lat0.value}, {alt0.value}) => ({lon.value}, {lat.value}, {alt.value})')
# (5.0, 0.0, 0.0) => (0.0, 555891.2675813203, -24270.736897208728)
