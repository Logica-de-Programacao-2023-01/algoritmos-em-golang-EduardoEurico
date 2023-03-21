package main

import "fmt"

func main() {
	var n int
	fmt.Print("Qual a base: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {

		fmt.Println(i * n)
	}
}
