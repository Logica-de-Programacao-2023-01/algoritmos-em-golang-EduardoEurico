package main

import "fmt"

func main() {

	var idade int
	var nome string
	var peso int

	fmt.Print("Qual é a sua idade? ")
	fmt.Scan(&idade)

	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&nome)

	fmt.Print("Qual é o seu peso? ")
	fmt.Scan(&peso)

	fmt.Print("Ola ", nome, " você atualmente tem ", idade, " anos e ", peso, "kg.")
}
