package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"

	alumno "./packages"
)

var alumnos = make(map[string]map[string]float64)

type Server struct {
}

func (this *Server) AgregarCalificacion(stu alumno.Alumno, reply *bool) error {

	materia := make(map[string]float64)
	materia[stu.Materia] = stu.Calificacion

	listSize := len(alumnos)

	// comprobemos que si existe el alumno
	flag := false
	for valor := range alumnos {
		if valor == stu.Nombre {
			flag = true
		}
	}

	if flag { // añadimos los datos al alumno existente en el map
		//fmt.Println("sí existe")
		alumnos[stu.Nombre][stu.Materia] = stu.Calificacion
		if len(alumnos[stu.Nombre]) > listSize {
			*reply = true
		}
		//fmt.Println("Número de alumnos:", len(alumnos[stu.Nombre]))
	} else { // añadimos los datos en un nuevo lugar dentro del map
		//fmt.Println("no existe")
		alumnos[stu.Nombre] = materia
		if len(alumnos) > listSize {
			*reply = true
		}
		//fmt.Println("Número de alumnos:", len(alumnos))
	}

	return nil // no hay error
}

func (this *Server) MostrarPromedioAlumno(nombre string, reply *float64) error {
	flag := false
	// comprobemos que la lista no esté vacía
	if len(alumnos) > 0 {
		// comprobemos que haya registros de ese alumno
		for valor := range alumnos {
			if valor == nombre {
				flag = true
			}
		}
		if flag {
			// obtenemos el promedio
			var sumatoria float64
			for _, cal := range alumnos[nombre] {
				sumatoria = sumatoria + cal
			}
			promedio := sumatoria / float64(len(alumnos[nombre]))
			*reply = promedio
		} else {
			return errors.New("No existe ese alumno")
		}
	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil
}

func (this *Server) MostrarPromedioGeneral(promedioGeneral float64, reply *float64) error {

	if len(alumnos) > 0 { // si la lista no está vacía
		var prom float64
		numAlumnos := float64(len(alumnos))

		for valor := range alumnos {
			var sumatoria float64
			for _, cal := range alumnos[valor] {
				sumatoria = sumatoria + cal
			}
			prom = sumatoria / float64(len(alumnos[valor]))
			promedioGeneral = promedioGeneral + prom
		}
		promedioGeneral = promedioGeneral / numAlumnos
		*reply = promedioGeneral
	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil // no hay error
}

func (this *Server) MostrarPromedioMateria(materia string, reply *float64) error {

	if len(alumnos) > 0 {
		// buscaremos la materia en cada alumno registrado
		flag := false
		var sumatoria float64
		contador := 0
		for valor := range alumnos {
			for mat, cal := range alumnos[valor] {
				if mat == materia {
					flag = true // siempre que se active la bandera verdaderamente, tomamos la calificación
				}
				if flag {
					sumatoria = sumatoria + cal
					contador = contador + 1
				}
			}
		}

		promedio := sumatoria / float64(contador)
		*reply = promedio

	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil // no hay error
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999") // arancamos el servidor
	if err != nil {
		fmt.Println(err)
	} // si hay error, lo imprimimos

	for {
		c, err := ln.Accept() // interceptamos al cliente
		if err != nil {
			fmt.Println(err)
			continue
		} // si hay error, lo imprimimos y continuamos

		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var waitKey string
	fmt.Scanln(&waitKey)
}
