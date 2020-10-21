/*
Crea un generador de número impares tomando de ejemplo generadorPares().
*/
package main

import "fmt"

func numimpar(a *int, b *int){

	for x:=*a; x<*b; x++{
		if x%2==1{
			fmt.Println(x)
		}
		
	}

}

func main(){
	var a,b int
	fmt.Println("Rango inicial: ")
	fmt.Scan(&a)
	fmt.Println("Rango final: ")
	fmt.Scan(&b)

	numimpar(&a, &b)

}