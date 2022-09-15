from ctypes import *

class CoordinateConverter(object):
	"""docstring for CoordinateConverter"""
	def __init__(self, lon, lat, alt):
		super(CoordinateConverter, self).__init__()
		self.dll = CDLL('../c_lib/libcoorconv.so')
		# 设置参数和返回值
		self.dll.coordinate_converter_new.restype = c_void_p
		self.dll.coordinate_convert.argtypes = [c_void_p, c_double, c_double, c_double, POINTER(c_double), POINTER(c_double), POINTER(c_double)]
		self.dll.coordinate_convert.restype = c_void_p
		self.dll.coordinate_converter_free.argtypes = [c_void_p]
		self.converter = self.dll.coordinate_converter_new(c_double(lon), c_double(lat), c_double(alt))

	def convert(self, lon, lat, alt):
		lon0 = c_double(lon)
		lat0 = c_double(lat)
		alt0 = c_double(alt)
		lon1 = c_double()
		lat1 = c_double()
		alt1 = c_double()
		self.dll.coordinate_convert(self.converter, lon0, lat0, alt0, byref(lon1), byref(lat1), byref(alt1))
		return (lon1.value, lat1.value, alt1.value)

	def __del__(self):
		self.dll.coordinate_converter_free(self.converter)

if __name__ == '__main__':
	converter = CoordinateConverter(0.0, 0.0, 0.0)
	res = converter.convert(5.0, 0.0, 0.0)
	print(f'res: {res}')
	# res: (0.0, 555891.2675813203, -24270.736897208728)
