package main

import "log"

/*
	Function structure in go

	func ( r receiver) Identifier(paramaters) (returns){code}
*/
/*
	Hello returns a greeting for provided name
*/
func Hello(name string) string {
	return "Hello " + name
}

/*
	HelloWithReceiver Is similar to Hello but no
		need for params since it is using the receiver
*/

func (n Name) HellWithReceiver() {
	log.Println("Heeeeeelllllooooo ", n.Name, " ", n.LastName)
}
