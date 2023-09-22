package ejercicios

import (
	"strconv"
)

func Ejercicio(texto string) (int, string) {
	entero, _ := strconv.Atoi(texto)
	if entero > 100 {
		return entero, "Es mayor a 100"
	} else {
		return entero, "Es menor a 100"
	}

}
