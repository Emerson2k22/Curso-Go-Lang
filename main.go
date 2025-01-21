package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Começando meus primeiros passos com a linguagem GoLang!")

//Tipos 

//bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

// string - sequência de bytes
    fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")

// inteiro
	fmt.Printf("Type: %T - Value: %v\n", 3, 3)
	

// float (float64/float32) - decimal
    fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)

	// Variáveis
	//var + nome da variável + tipo
	var nome string
	nome = "bento"
	fmt.Println(nome)

	// var int
	var idade int
	idade = 4
	fmt.Println(idade)
	// ou para mais de uma variável
	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	// variável bool
	var d = true
	fmt.Println(d)

    // variável abreviada
	f := "apple"
	fmt.Println(f)

	// constantes
	const idadeBento = 24
	fmt.Println(idadeBento)

	// funções
	fmt.Println(soma(42, 13))
}

func soma (x int, y int) int {
	return x + y
}