package main

import "fmt"

func main() {
	slice := []int{5, 3, 7, 4, 1}
	fmt.Println(ordenar(slice))
}

func ordenar(sliceNumeros []int) []int {
	acumuladorSlice := make([]int, len(sliceNumeros))
	acumuladorNumero := sliceNumeros[0]
	acumuladorIndice := 0
	for j := 0; j < len(acumuladorSlice); j++ {
		for i := 0; i <= len(sliceNumeros)-1; i++ {
			if sliceNumeros[i] < acumuladorNumero {
				acumuladorNumero = sliceNumeros[i]
				acumuladorIndice = i
			}
		}
		acumuladorSlice[j] = acumuladorNumero
		sliceNumeros = append(sliceNumeros[:acumuladorIndice], sliceNumeros[acumuladorIndice+1:]...)
		acumuladorIndice = 0
		if len(sliceNumeros) == 0 {
			return acumuladorSlice
		} else {
			acumuladorNumero = sliceNumeros[0]
		}
	}
	return acumuladorSlice
}
