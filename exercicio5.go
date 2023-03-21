package main

import "fmt"

func main() {

	var real float64
	fmt.Print("Escolha um valor de real ")
	fmt.Scan(&real)

	var dolar float64 = (real / 5.24)
	fmt.Print("O valor em dolar é ", dolar)
}
