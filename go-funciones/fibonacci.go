package main

import "fmt"

func Fibonacci(number int64) int64 {
	if number == 0 || number == 1 {
		return number
	}
	return Fibonacci(number-2) + Fibonacci(number-1)
}

func main() {
	fmt.Println(Fibonacci(65))
}
