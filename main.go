package main

import "fmt"
//go:generate stringer -type=Pill
func main()  {

fmt.Println("hello raj")
fmt.Println(Fib(5))
var x Pill = Aspirin
x = Paracetamol
fmt.Println(x.String())
}

func Fib(a int) int  {
	if a <= 1 {
		return a
	}
	return Fib(a-1)+Fib(a-2)
}
