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
}