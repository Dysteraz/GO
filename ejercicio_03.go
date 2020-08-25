package main

import "fmt"

func asd2() {
	fmt.Println(esPalindromo("martin"))
}

func esPalindromo(palabra string) bool {
	largoIndexPalabra := len(palabra) - 1
	for i := 0; i <= largoIndexPalabra; i++ {
		if palabra[i] == palabra[largoIndexPalabra] {
			largoIndexPalabra--
		} else {
			return false
		}
	}
	return true

}
