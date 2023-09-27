package main

import (
	e "D/Workspace/Go/godesde0/ejer_interfaces"
	"D/Workspace/Go/godesde0/modelos"
)

func main() {
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)
	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
	//users.AltaUsuario()
	//mapas.MostrarMapas()
	//arreglos_slices.Capacidad()
	//arreglos_slices.MuestroSlice()
	//funciones.Exponencia(2)
	//funciones.LlamarClosure()
	//funciones.Calculos()

	//files.LeoArchivo()

	//files.SumaTabla()

	//iteraciones.Iterar()

	//teclado.IngresosNumeros()

	/* numero, texto := ejercicios.Ejercicio("fffff")
	fmt.Println(numero, texto) */

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
