package main

type MyError struct {
	str string
}

func (e *MyError) Error() string {
	return "123123123123"
}

func returnsError() error {
	var p *MyError = nil
	return p
}

func main() {
	err := returnsError()
	println(err)
	if err != nil {
		// 不为空，这里会打印
		println("err not nil")
		println("err: " + err.Error())
	}
}
