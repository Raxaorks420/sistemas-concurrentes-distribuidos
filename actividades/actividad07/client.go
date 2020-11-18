package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	proceso "./packages"
)

var bandera = false
var proc proceso.Proceso

func client() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(bandera)
	if err != nil {
		fmt.Println(err)
	} // avisamos al server del estado de la bandera

	if !bandera { // si es falsa, se recibe un proceso
		err = gob.NewDecoder(c).Decode(&proc)
		if err != nil {
			fmt.Println(err)
			return
		}
		bandera = true
	}

	c.Close()

}

func main() {
	go client()

	go func() {
		for {
			fmt.Println(proc.Id, " : ", proc.Contador)
			proc.Contador++
			fmt.Println("-------------------------------------------------")
			time.Sleep(time.Millisecond * 500)
		}
	}()

	var input string
	fmt.Scanln(&input)

	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(bandera)
	if err != nil {
		fmt.Println(err)
	}

	err = gob.NewEncoder(c).Encode(proc)
	if err != nil {
		fmt.Println(err)
	}
	bandera = false

	fmt.Println("Cliente Terminado\n")
}
