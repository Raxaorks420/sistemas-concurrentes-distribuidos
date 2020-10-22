package multimedia

import "fmt"

type ContenidoWeb struct {
	Juas []Multimedia
}

type Multimedia interface {
	Mostrar()
}

func (cw *ContenidoWeb) Mostrar() {
	for _, c := range cw.Juas {
		c.Mostrar()
		fmt.Println("")
	}
}

type Imagen struct { // estructura Imagen
	Titulo  string
	Formato string
	Canal   string
}

func (i *Imagen) Mostrar() { // método Mostrar de Imagen
	fmt.Println(".: IMAGEN :.")
	fmt.Println(" Titulo:", i.Titulo)
	fmt.Println(" Formato:", i.Formato)
	fmt.Println(" Canal:", i.Canal)
}

type Audio struct { // estructura Audio
	Titulo   string
	Formato  string
	Duracion int
}

func (a *Audio) Mostrar() { // método Mostrar de Audio
	fmt.Println(".: AUDIO :.")
	fmt.Println(" Titulo:", a.Titulo)
	fmt.Println(" Formato:", a.Formato)
	fmt.Println(" Duración:", a.Duracion, "segundos")
}

type Video struct { // estructura Video
	Titulo  string
	Formato string
	Frames  int
}

func (v *Video) Mostrar() { // método Mostrar de Video
	fmt.Println(".: VIDEO :.")
	fmt.Println(" Titulo:", v.Titulo)
	fmt.Println(" Formato:", v.Formato)
	fmt.Println(" Frames:", v.Frames)
}
