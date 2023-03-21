package main

import "fmt"

func main() {

	var numero1 int
	fmt.Print("Digite o 1º número: ")
	fmt.Scan(&numero1)

	var numero2 int
	fmt.Print("Digite o 2º número: ")
	fmt.Scan(&numero2)

	var numero3 int
	fmt.Print("digite o 3º número: ")
	fmt.Scan(&numero3)

	if numero1 < numero2 {
		fmt.Print("o menor numero é ", numero1)
	} else if numero2 > numero3 {
		fmt.Print("o menor numero é ", numero2)
	} else {
		fmt.Print("O menor número é ", numero3)
	}
	if numero1 == numero2 {
		fmt.Print(" ,e os numeros são iguais")
	} else if numero1 == numero3 {
		fmt.Print(" ,e os numeros são iguais")
	}
}
