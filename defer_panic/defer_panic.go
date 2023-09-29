package defer_panic

import (
	"fmt"
	"log"
)

// El defer nos permite configurar cual es la ultima linea que se va a ejecutar cuando finalice la ejecucion de la funcion
func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

// El panic me permite cortar con la ejecucion del sistema
// El recover se usa con defer si o si. Cuando el recover detecta que hubo un panic el valor que retorna no es nil
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
