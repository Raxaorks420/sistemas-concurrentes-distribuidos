package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println("\n\t STRINGS/ARCHIVOS/SORT")
	fmt.Print(" -> Cuántos strings vas a introducir? ")
	var n int
	fmt.Scanln(&n) // pidiendo el número de strings

	s := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		s = append(s, line)
	} // llenando el slice n-veces

	// creando archivos
	file1, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()

	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	fmt.Println("Slice desordenado :", s)
	sort.Strings(s)
	fmt.Println("Slice ascendente  :", s) // slice ordenado ascendentemente
	for _, v := range s {
		file1.WriteString(v + "\n") // escribiendo el archivo ascendente.txt
	}

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Slice descendente :", s) // slice ordenado descendentemente
	for _, v := range s {
		file2.WriteString(v + "\n") // escribiendo el archivo descendente.txt
	}
}
