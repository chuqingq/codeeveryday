package main

import ()

func main() {
	myarray := new([10]int)
	println("myarray: cap(myarray):", cap(myarray), " len(myarray):", len(myarray))
	myalice := myarray[2:4]
	println("myalice: cap(myalice):", cap(myalice), " len(myalice):", len(myalice))
}
