package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() string {
	var num int
	var err error
	var texto string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}
	return texto
}
