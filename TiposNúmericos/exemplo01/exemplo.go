package main

import "fmt"

func main() {
	// Tipo int se usa para números com sinais e uint para números sem sinais.
	var quantidade uint = 34
	var temperatura int = -34
	fmt.Printf("Quantidade: %d\nTemperatura: %d\n", quantidade, temperatura)

	var peso float64 = 1.5
	fmt.Printf("Peso: %f\n", peso)
} 