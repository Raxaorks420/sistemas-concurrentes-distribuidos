package main

import (
	"fmt"
	"os"
	"time"
)

// variables globales
var mostrar bool = false
var lista = []Proceso{}
var contador int = 0

type Proceso struct {
	Id     uint64
	Estado string
	quit   chan bool
}

func (p *Proceso) start() {
	p.quit = make(chan bool)
	aidi := p.Id
	go func() {
		i := uint64(0)
		for {

			select {
			case <-p.quit:
				return
			default:
				if mostrar {
					fmt.Println(aidi, ":", i)
				}
				i = i + 1
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func (p *Proceso) stop() {
	p.quit <- true
}

func agregarProceso() {
	var canal chan bool
	p := Proceso{uint64(contador), "Ejecutando", canal} // creando un nuevo proceso
	contador = contador + 1

	lista = append(lista, p) // lo añadimos a un slice
	pos := len(lista) - 1
	lista[pos].start() // iniciamos el proceso recién añadido, el cual es el último del slice
	fmt.Println("\t [OK] Proceso creado e iniciado. Regresando al menú . . .")
}

func mostrarProceso() {

	if len(lista) == 0 {
		fmt.Println("\t [!] Lista vacía! Regresando al menú . . .")
	} else {
		mostrar = true

		var line string
		fmt.Scanln(&line)

		mostrar = false
	}

}

func terminaProceso() {
	// pedimos el id del proceso y lo terminamos
	if len(lista) == 0 {
		fmt.Println("\t [!] Lista vacía! Regresando al menú . . .")
	} else {
		// pedimos el id del proceso a eliminar
		fmt.Print("ID: ")
		var res uint64
		fmt.Scanln(&res)

		flag := false

		for index, valor := range lista {
			if res == valor.Id {
				valor.stop()
				valor.Estado = "Terminado"
				// eliminamos el proceso de la lista
				lista = append(lista[:index], lista[index+1:]...)
				flag = true
				fmt.Println("\t [OK] Proceso eliminado y rutina terminada. Regresando al menú . . .")
				break
			}
		}
		if !flag {
			fmt.Println("\t [!] El ID que ingresaste no existe. Regresando al menú . . .")
		}
	}
}

func menu() {
	for {
		fmt.Println("\n\tACTIVIDAD 06 - Goroutines")
		fmt.Println(" 1) Agregar proceso")
		fmt.Println(" 2) Mostrar procesos")
		fmt.Println(" 3) Terminar proceso")
		fmt.Println(" 4) SALIR")
		fmt.Print(" --> Opción: ")

		var opc int
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			agregarProceso()
		case 2:
			mostrarProceso()
		case 3:
			terminaProceso()
		case 4:
			fmt.Println("\t [!] Saliendo del programa . . .")
			fmt.Scanln()
			os.Exit(2)
		default:
			fmt.Println("\t [!] Opción incorrecta, intenta de nuevo . . .")
			fmt.Scanln()
		}
	}
}

func main() {
	menu()
}
