package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	proceso "./packages"
)

var lista_procesos []proceso.Proceso
var status bool = false

func creaLista() []proceso.Proceso {
	lista := make([]proceso.Proceso, 0)
	for i := 0; i < 5; i++ {
		p := proceso.Proceso{Id: uint64(i), Contador: uint64(0)}
		lista = append(lista, p)
	}
	return lista
}

func arrancaProcesos() {
	for {
		for index, value := range lista_procesos {
			fmt.Println(value.Id, " : ", value.Contador)
			lista_procesos[index].Contador++
		}
		fmt.Println("-------------------------------------------------")
		time.Sleep(time.Millisecond * 500)
	}
}

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	} // arrancamos el servidor

	lista_procesos = creaLista()
	go arrancaProcesos()

	for {
		c, err := s.Accept() // nos comunicamos con algún cliente
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c) // lanzamos función concurrente
	}
}

func handleClient(c net.Conn) {

	if !status {
		err := gob.NewDecoder(c).Decode(&status)
		if err != nil {
			fmt.Println(err)
			return
		} // recibimos estado del cliente
	}

	if status {
		// el cliente se desconecta y retorna un proceso
		var p proceso.Proceso
		err := gob.NewDecoder(c).Decode(&p)
		if err != nil {
			fmt.Println(err)
			fmt.Println("asdasd")
		}
		lista_procesos = append(lista_procesos, p)
		status = false

	} else { // el cliente extrae un proceso del servidor
		proc := lista_procesos[len(lista_procesos)-1] // tomamos el último proceso de la lista
		err := gob.NewEncoder(c).Encode(proc)         // lo codificamos y enivamos
		if err != nil {
			fmt.Println(err)
		} else {
			if len(lista_procesos) > 0 {
				lista_procesos = lista_procesos[:len(lista_procesos)-1]
			} // si envía exitosamente, lo eliminamos de la lista
		}
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
	fmt.Println("Servidor Terminado\n")
}
