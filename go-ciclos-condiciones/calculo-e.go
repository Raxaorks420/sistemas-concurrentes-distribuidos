package main

import "fmt"

func factorial(n int) float64 {
	// toma un número entero como parámetro, siendo este n, y calcula su factorial
	var valor float64 = 1
	if n < 0 {
		fmt.Println("No podemos calcular factorial de un número negativo")
	} else {
		for i := 1; i <= n; i++ {
			valor *= float64(i)
		}
	}
	return valor
}

func main() {
	/* pedimos al usuario el valor de n*/
	var number int
	fmt.Scan(&number)

	var euler float64
	for i := 0; i <= number; i++ {
		euler = euler + (1 / factorial(i))
	}

	fmt.Println(euler)
}
