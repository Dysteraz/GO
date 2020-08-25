package main

import "fmt"

func qwert() {
	fmt.Println(enumerar("pan"))
}

func enumerar(palabras ...string) string {
	largoDeArray := len(palabras)
	acumulador := ""
	for i := 0; i <= largoDeArray-1; i++ {
		if i == 0 {
			acumulador = acumulador + palabras[i]
		} else if i != largoDeArray-1 {
			acumulador = acumulador + ", " + palabras[i]
		} else {
			acumulador = acumulador + " y " + palabras[i]
		}
	}
	return acumulador
}
