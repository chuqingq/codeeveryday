def my_print(int_arg):
	local a
	if a == None:
		a = 0 
	a = a + int_arg
	print a

def test(myfun, int_arg):
	myfun(int_arg)

test(my_print, 1)
test(my_print, 2)

