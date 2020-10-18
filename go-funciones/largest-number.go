package main

import "fmt"

func largestNumber(args ...int) int {
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func main() {
	fmt.Println(largestNumber(13, 12, 88, 190, 210, 14, 5, 7))
}
