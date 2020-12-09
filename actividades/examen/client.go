package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var writeStr, readStr = make([]byte, 2048), make([]byte, 2048)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "8000"
		remote = host + ":" + port
		reader = bufio.NewReader(os.Stdin)
	)

	con, err := net.Dial("tcp", remote)
	defer con.Close()

	if err != nil {
		fmt.Println("El servidor no se ha encontrado.\n")
		os.Exit(-1)
	}
	fmt.Println("Estableciendo conexión.\n")
	fmt.Printf("-> Nickname: ")
	fmt.Scanf("%s", &writeStr)
	in, err := con.Write([]byte(writeStr))
	if err != nil {
		fmt.Printf("Error al conectarse al servidor: %d\n", in)
		os.Exit(0)
	}
	fmt.Println("\nConexión exitosa!")

	go read(con)

	// aquí hay que meter el menú
	fmt.Println("\n\t CHATROOM - Menú principal")
	fmt.Println("1) Enviar mensaje *")
	fmt.Println("2) Mostrar la conversación completa")
	fmt.Println("3) Salir")
	fmt.Println("-> Teclea el número de la opción que desees y presiona enter.")
	fmt.Println(" * si eliges la opción 1, teclea el mensaje que quieres enviar después de presionar enter.\n")

	for {
		writeStr, _, _ = reader.ReadLine()

		if string(writeStr) == "1" {
			writeStr, _, _ = reader.ReadLine()
			mensaje := string(writeStr)
			in, err := con.Write([]byte(mensaje))
			if err != nil {
				fmt.Printf("Error al conectarse al servidor: %d\n", in)
				os.Exit(0)
			}

		} else if string(writeStr) == "2" {
			mensaje := "CMD:FULL"
			in, err := con.Write([]byte(mensaje))
			if err != nil {
				fmt.Printf("Error al conectarse al servidor: %d\n", in)
				os.Exit(0)
			}
		} else if string(writeStr) == "3" {
			fmt.Println("Saliendo del sistema. Desconectando al cliente...\n")
			os.Exit(1)
		}

	}
}

func read(conn net.Conn) {
	for {
		length, err := conn.Read(readStr)
		if err != nil {
			fmt.Printf("Error al leer del servidor. Error:%s\n", err)
			os.Exit(0)
		}
		fmt.Println(" ----------------------------------\n[INCOMING] \n" + string(readStr[:length]) + "\n ------------------------------------ ")
	}
}
