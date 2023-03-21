package main

import "fmt"

func main() {

	var base int
	fmt.Print("Qual a base: ")
	fmt.Scan(&base)

	var altura int
	fmt.Print("Qual a altura: ")
	fmt.Scan(&altura)

	var area_retangulo int = base * altura

	fmt.Print("A área do retangulo de base ", base, " e altura ", altura, " é ", area_retangulo, "m².")

}
