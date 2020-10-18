package main

import "fmt"

func intercambia(n1 *int, n2 *int) {
	aux1 := *n1
	aux2 := *n2
	*n1 = aux2
	*n2 = aux1
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	intercambia(&a, &b)
	fmt.Println(a, b)
}
