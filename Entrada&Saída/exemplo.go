package main

import "fmt"

func main() {
	fmt.Println("Digite o seu nome:")
	var nome string
	fmt.Scanf("%s", &nome)
	fmt.Printf("Seja bem-vindo(a), %s!", nome)
}