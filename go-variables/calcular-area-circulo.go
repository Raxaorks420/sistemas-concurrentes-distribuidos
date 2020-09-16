package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\nÁREA DE UN CÍRCULO")

	var radio float64
	fmt.Print("\n -> Ingresa el tamaño del radio: ")
	fmt.Scan(&radio)

	area := math.Pi * (radio * radio)
	fmt.Println("\n ---> ÁREA: ", area)
}
