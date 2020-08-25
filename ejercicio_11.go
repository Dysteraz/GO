package main

import "fmt"

func main() {
	fmt.Println(disanciaDeLevenstein("coma", "cola"))
}

func disanciaDeLevenstein(palabra1 string, palabra2 string) int {
	if palabra1 == palabra2 {
		return 0
	}
	contador := 0
	indexMaximo := len(palabra1) - 1
	if len(palabra1) > len(palabra2) {
		contador = len(palabra1) - len(palabra2)
		indexMaximo = len(palabra2) - 1
	} else if len(palabra2) > len(palabra1) {
		contador = len(palabra2) - len(palabra1)
		indexMaximo = len(palabra1) - 1
	}
	for i := 0; i <= indexMaximo; i++ {
		if palabra1[i] != palabra2[i] {
			contador++
		}
	}
	return contador
}
