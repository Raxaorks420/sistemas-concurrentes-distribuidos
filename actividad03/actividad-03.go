package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("\nSISTEMAS CONCURRENTES Y DISTRIBUIDOS - Actividad 03: Git branch")

	var line string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n -> Ingresa una cadena de texto: ")
	scanner.Scan()
	line = scanner.Text()

	if isPalindrome(line) {
		fmt.Println(" ---> Sí es un palíndromo! XD")
	} else {
		fmt.Println(" ---> No es un palíndromo! xC")
	}

	fmt.Println()
}
