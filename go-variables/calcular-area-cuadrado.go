package main

import "fmt"

func main() {
	fmt.Println("\nÁREA DE UN CUADRADO")

	var lado float64

	fmt.Print(" -> Ingresa el tamaño del lado del cuadrado: ")
	fmt.Scan(&lado)

	area := lado * lado
	fmt.Println("---> ÁREA: ", area)

}
