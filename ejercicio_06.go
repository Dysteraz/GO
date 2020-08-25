package main

import "fmt"

func qwer() {
	fmt.Println(reemplazarPalabra("Hola soy homero simpson", "soy", "ROBERTOVICH"))
}

func reemplazarPalabra(frase string, palabra1 string, palabra2 string) string {
	largoPalabra := len(palabra1)
	for i := 0; i <= len(frase)-len(palabra1); i++ {
		if frase[i:largoPalabra] == palabra1[:] {
			sliceFrase1 := frase[0:i]
			sliceFrase2 := frase[largoPalabra:len(frase)]
			return sliceFrase1 + palabra2 + sliceFrase2
		}
		largoPalabra++
	}
	return "Palabra no encontrada"
}
