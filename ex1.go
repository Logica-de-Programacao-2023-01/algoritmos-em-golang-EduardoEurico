package main

import "fmt"

func main() {
	var numero1 int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&numero1)

	var numero2 int
	fmt.Print("Digite outro numero: ")
	fmt.Scan(&numero2)

	if numero1 < numero2 {
		fmt.Println(numero1, " é menor que ", numero2)
	} else {
		fmt.Println(numero1, " é maior que ", numero2)
	}
}
