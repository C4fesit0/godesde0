package main

import (
	"D/Workspace/Go/godesde0/ejercicios"
	"fmt"
)

func main() {
	numero, texto := ejercicios.Ejercicio("20")
	fmt.Println(numero, texto)

	/* estado, texto := variables.ContvierteaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */

	/* if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es ", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Println("&s \n", os)
	}*/

}
