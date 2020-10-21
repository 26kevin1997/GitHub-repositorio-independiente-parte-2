/*  
Implementar la secuencia de Fibonacci de manera recursiva
*/
package main

import "fmt"

func finobacci(a *int, b *int){
	var s []int

	s = append(s, *a)
	s = append(s, *b)
	fmt.Println()
	fmt.Println(s[0])
	fmt.Println(s[1])
	for x:=0; x<20; x++{
		res:=s[x]+s[x+1]
		s=append (s, res)
		fmt.Println(res)
	}

}

func main(){
	var a,b int
	fmt.Println("Numero 1: ")
	fmt.Scan(&a)
	fmt.Println("Numero 2: ")
	fmt.Scan(&b)

	finobacci(&a, &b)

}