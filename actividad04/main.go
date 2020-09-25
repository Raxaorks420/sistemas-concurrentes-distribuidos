package main

import "fmt"

func main() {
	var base1, base2, altura float64
	fmt.Println("\n Calcular el área de un TRAPECIO")

	fmt.Print(" -> Medida de la base menor : ")
	fmt.Scan(&base1)
	fmt.Print(" -> Medida de la base mayor : ")
	fmt.Scan(&base2)
	fmt.Print(" -> Medida de la altura : ")
	fmt.Scan(&altura)

	area := ((base1 + base2) * altura) / 2
	fmt.Println("\n ---> ÁREA : ", area)

}
