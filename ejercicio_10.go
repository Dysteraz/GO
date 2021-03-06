package main

import "fmt"

func htr() {
	fmt.Println(tieneParentesisBalanceados("()()()"))
}

func tieneParentesisBalanceados(unStrinConParentesis string) bool {
	contador := 0
	parentesisCerrar := ")"
	parentesisAbrir := "("
	for i := 0; i < len(unStrinConParentesis); i++ {
		if contador < 0 {
			return false
		} else if unStrinConParentesis[i] == parentesisCerrar[0] {
			contador--
		} else if unStrinConParentesis[i] == parentesisAbrir[0] {
			contador++
		}
	}
	return contador == 0
}
