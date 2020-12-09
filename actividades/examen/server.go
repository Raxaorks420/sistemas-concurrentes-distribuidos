package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var cliQue []net.Conn
var conversation []string

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nServidor finalizado.")
		SaveConvToFile()
		os.Exit(0)
		// aquí guardamos la conversación en un archivo
	}()
}

func SaveConvToFile() {
	file, err := os.Create("conversation.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // cerramos el archivo al terminar la función

	file.WriteString(GetConv())
}

func GetConv() string {
	var newstr string
	for _, value := range conversation {
		newstr += value + "\n"
	}
	return newstr
}

// Reenvía el mensaje a los demás clientes conectados
func notify(conn net.Conn, msg string) {
	for _, con := range cliQue { // recorremos la cola de clientes
		if con.RemoteAddr() != conn.RemoteAddr() { // si la dirección es distinta a la emisora
			con.Write([]byte(msg)) // lanzamos el mensaje
		}
	}
}

// Desconexión de los clientes
func disconnect(conn net.Conn, name string) {
	for index, con := range cliQue {
		if con.RemoteAddr() == conn.RemoteAddr() {
			disMsg := name + " se ha desconectado.\n"
			//fmt.Println(disMsg)
			cliQue = append(cliQue[:index], cliQue[index+1:]...)
			notify(conn, disMsg)
		}
	}
}

func main() {
	var (
		host   = "127.0.0.1"
		port   = "8000"
		remote = host + ":" + port
		data   = make([]byte, 2048)
	)
	fmt.Println("Cargando servidor... (Ctrl-C para terminar)\n")

	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	SetupCloseHandler()

	for {
		var res string
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error al conectar con cliente: ", err.Error())
			os.Exit(0)
		}
		cliQue = append(cliQue, conn)

		go func(con net.Conn) {
			fmt.Println("Nueva conexión: ", con.RemoteAddr())

			// Get client's name
			length, err := con.Read(data)
			if err != nil {
				fmt.Printf("El cliente %v se ha desconectado.\n", con.RemoteAddr())
				con.Close()
				disconnect(con, con.RemoteAddr().String())
				return
			}
			name := string(data[:length])
			comeStr := name + " se ha conectado.\n"
			notify(con, comeStr)
			fmt.Println("\n" + comeStr)

			// Begin recieve message from client
			fmt.Println()
			for {
				length, err := con.Read(data)
				if err != nil {
					fmt.Printf("\n%s se ha desconectado.\n", name)
					con.Close()
					disconnect(con, name)
					return
				}
				res = string(data[:length])
				if res == "CMD:FULL" {
					// aquí retornamos una lista con la conversación y se manda
					listaStr := GetConv()
					con.Write([]byte(listaStr))
				} else {
					sprdMsg := "-> " + name + " : " + res
					fmt.Println(sprdMsg)
					res = sprdMsg
					con.Write([]byte(res))
					notify(con, sprdMsg)

					conversation = append(conversation, res)
				}

			}
		}(conn)
	}

}
