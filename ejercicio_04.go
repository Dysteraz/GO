package main

func asd1() {
	println(contarLetra("anana", "a"))
}

func contarLetra(palabra string, letraAContar string) int {
	contador := 0
	for i := 0; i <= len(palabra)-1; i++ {
		if palabra[i] == letraAContar[0] {
			contador++
		}
	}
	return contador
}
