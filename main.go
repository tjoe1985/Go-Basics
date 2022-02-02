package main

import "log"

func main() {
	//call a variadic function and print passed params and the type
	Variadic(1, 2, 3, 4, 5, 6, 7)

	//pass slice as variadic parameter.
	x := []int{8, 9, 10, 11, 12, 13, 14}
	Variadic(x...)

	// calling regular function Hello
	s := Hello("Joel")
	log.Println(s)

	// calling function with receiver
	subject := Name{
		Name:     "Kate",
		LastName: "Smith",
	}
	subject.HellWithReceiver()
}
