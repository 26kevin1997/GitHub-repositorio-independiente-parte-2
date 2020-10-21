/*
Crea un programa que pida dos enteros (a y b)
y llame a la función intercambia(&a, &b),
la cual intercambia el valor de a y b; usando punteros.
*/
package main

import "fmt"

func intercambia(a *int, b *int) {
	var aa, bb int
	aa = *b
	bb = *a

	fmt.Println("Numero A: ", aa)
	fmt.Println("Numero B: ", bb)

}

func main() {
	var a, b int
	fmt.Println("Numero A: ")
	fmt.Scan(&a)
	fmt.Println("Numero B: ")
	fmt.Scan(&b)

	intercambia(&a, &b)

}
