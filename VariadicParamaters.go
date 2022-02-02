package main

import "fmt"

/*
	Variadic will print passed parameters and type.
*/
func Variadic(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	for i, val := range x {
		fmt.Println("Value of x at Index: ", i, " is: ", val)
	}
}
