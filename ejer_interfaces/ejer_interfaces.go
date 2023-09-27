package ejer_interfaces

import (
	"D/Workspace/Go/godesde0/interfaces"
	"fmt"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando", hu.Sexo())
}

func SeresVivos() {

}
