package main

import "fmt"

func main() {
	name := "Jeff"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)
}

func changeMe(z string) {
	fmt.Println(z)
	z = "Rusty"
	fmt.Println(z)
}
