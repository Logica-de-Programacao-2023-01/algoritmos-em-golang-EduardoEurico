package main

import "fmt"

func main() {

	var numero1 int
	fmt.Print("Escolha um número")
	fmt.Scan(&numero1)

	var numero2 int
	fmt.Print("Escolha um número")
	fmt.Scan(&numero2)

	var numero3 int
	fmt.Print("Escolha um número")
	fmt.Scan(&numero3)

	var media int = (numero1 + numero2 + numero3) / 3

	fmt.Print("a media dos números é ", media, ".")

}
