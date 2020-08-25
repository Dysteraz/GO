package main

func qwe() {
	println(contarPalabra("una", "habia una vez una moto en una casa"))
}

func contarPalabra(palabra string, frase string) int {
	contador := 0
	largoPalabra := len(palabra)
	for i := 0; i <= len(frase)-len(palabra)-1; i++ {
		if frase[i:largoPalabra] == palabra[:] {
			contador++
		}
		largoPalabra++
	}
	return contador
}
