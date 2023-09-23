package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio2() {
	var num int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un numero")
	for {
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto ingreselo nuevamente")
			} else {
				break
			}
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(num, "x", i+1, "=", num*(i+1))
	}

}
