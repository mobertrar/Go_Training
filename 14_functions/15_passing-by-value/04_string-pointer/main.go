package main

import "fmt"

func main() {
	name := "rob"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name)
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rusty"
	fmt.Println(z)
	fmt.Println(*z)
}
