package main

import "fmt"

func main() {

	var base int
	fmt.Print("Qual a base: ")
	fmt.Scan(&base)

	var altura int
	fmt.Print("Qual a altura: ")
	fmt.Scan(&altura)

	var profundidade int
	fmt.Print("Qual a profundidade: ")
	fmt.Scan(&profundidade)

	var volume_caixa int = base * altura * profundidade

	fmt.Print("O volume da caixa de ", base, " e altura ", altura, " e ", " profundidade ", profundidade, " é ", volume_caixa, "m³.")

}
