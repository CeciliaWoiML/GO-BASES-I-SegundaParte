package main

import "fmt"

// fmt.Println("Ingrese una palabra: ")
//	var palabra string
//	fmt.Scanln((&palabra))

func main() {

	palabra := "Cecilia"

	fmt.Println("La cantidad de letras de la palabra es: ", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Println("Letra número ", i, ": ", string(palabra[i]))
	}	
}