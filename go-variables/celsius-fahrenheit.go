package main

import "fmt"

func main() {
	fmt.Println("\nConversor de grados CELSIUS a FAHRENHEIT")

	var gradcel float64
	fmt.Print("\n -> Ingresa los grados en CELSIUS: ")
	fmt.Scan(&gradcel)

	gradfahren := (gradcel * 9 / 5) + 32
	fmt.Println("\n ---> Grados FAHRENHEIT:", gradfahren)
}
