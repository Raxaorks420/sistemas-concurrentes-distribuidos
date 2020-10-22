package main

import (
	"bufio"
	"fmt"
	"os"

	"./multimedia"
)

func muestraMultimedia(multi ...multimedia.Multimedia) {
	for _, m := range multi {
		m.Mostrar()
		fmt.Println("")
	}
}

func main() {
	s := []multimedia.Multimedia{}
	// menú principal
	for {
		fmt.Println("\n\n\t ESTRUCTURAS / INTERFACES")
		fmt.Println("   1) Capturar Imagen")
		fmt.Println("   2) Capturar Audio")
		fmt.Println("   3) Capturar Video")
		fmt.Println("   4) Mostrar multimedia")
		fmt.Println("   5) SALIR")
		fmt.Print("\n   -> Opción: ")
		var opc int
		fmt.Scanln(&opc)

		switch opc {
		case 1:

			var _titulo, _formato, _canal string
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("\n Captura los datos de una imagen")
			fmt.Print("- Título  : ")
			scanner.Scan()
			_titulo = scanner.Text()
			fmt.Print("- Formato : ")
			scanner.Scan()
			_formato = scanner.Text()
			fmt.Print("- Canal   : ")
			scanner.Scan()
			_canal = scanner.Text()
			img := multimedia.Imagen{Titulo: _titulo, Formato: _formato, Canal: _canal}
			s = append(s, &img)

		case 2:
			var _titulo, _formato string
			var _duracion int
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("\n Captura los datos de un audio")
			fmt.Print("- Título   : ")
			scanner.Scan()
			_titulo = scanner.Text()
			fmt.Print("- Formato  : ")
			scanner.Scan()
			_formato = scanner.Text()
			fmt.Print("- Duración : ")
			fmt.Scanln(&_duracion)
			audio := multimedia.Audio{Titulo: _titulo, Formato: _formato, Duracion: _duracion}
			s = append(s, &audio)
		case 3:
			var _titulo, _formato string
			var _frames int
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("\n Captura los datos de un video")
			fmt.Print("- Título  : ")
			scanner.Scan()
			_titulo = scanner.Text()
			fmt.Print("- Formato : ")
			scanner.Scan()
			_formato = scanner.Text()
			fmt.Print("- Frames  : ")
			fmt.Scanln(&_frames)
			video := multimedia.Video{Titulo: _titulo, Formato: _formato, Frames: _frames}
			s = append(s, &video)
		case 4:
			fmt.Println(" Mostrando todos los datos capturados ")
			a := multimedia.ContenidoWeb{
				Juas: s,
			}
			a.Mostrar()
			fmt.Scanln()
		case 5:
			os.Exit(2)
		default:
			fmt.Println(" [!] Opción incorrecta, intenta de nuevo...")
		}
	}
}
