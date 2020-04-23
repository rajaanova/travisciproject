package main

import "fmt"

func main()  {

fmt.Println("hello raj")
fmt.Println(Fib(5))
}

func Fib(a int) int  {
	if a <= 1 {
		return a
	}
	return Fib(a-1)+Fib(a-2)
}
