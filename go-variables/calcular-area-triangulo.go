package main

import "fmt"

func main() {
	fmt.Println("\nÁREA DE UN TRIÁNGULO")

	var base, altura float64

	fmt.Print("\n -> Ingresa el tamaño de la base: ")
	fmt.Scan(&base)

	fmt.Print(" -> Ingresa el tamaño de la altura: ")
	fmt.Scan(&altura)

	area := (base * altura) / 2
	fmt.Println("\n ---> ÁREA: ", area)

}
