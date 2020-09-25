package main

import "fmt"

func main() {

	var day int
	fmt.Scan(&day)

	var month int
	fmt.Scan(&month)

	var signo string

	switch month {
	case 12:
		if day < 22 {
			signo = "sagitario"
		} else {
			signo = "capricornio"
		}
	case 11:
		if day < 22 {
			signo = "escorpio"
		} else {
			signo = "sagitario"
		}
	case 10:
		if day < 23 {
			signo = "libra"
		} else {
			signo = "escorpio"
		}
	case 9:
		if day < 23 {
			signo = "virgo"
		} else {
			signo = "libra"
		}
	case 8:
		if day < 23 {
			signo = "leo"
		} else {
			signo = "virgo"
		}
	case 7:
		if day < 23 {
			signo = "cancer"
		} else {
			signo = "leo"
		}
	case 6:
		if day < 21 {
			signo = "geminis"
		} else {
			signo = "cancer"
		}
	case 5:
		if day < 21 {
			signo = "tauro"
		} else {
			signo = "geminis"
		}
	case 4:
		if day < 20 {
			signo = "aries"
		} else {
			signo = "tauro"
		}
	case 3:
		if day < 21 {
			signo = "piscis"
		} else {
			signo = "aries"
		}
	case 2:
		if day < 19 {
			signo = "acuario"
		} else {
			signo = "piscis"
		}
	case 1:
		if day < 20 {
			signo = "capricornio"
		} else {
			signo = "acuario"
		}
	default:
		fmt.Println("Intenta de nuevo . . .")
	}

	fmt.Println(signo)
}
