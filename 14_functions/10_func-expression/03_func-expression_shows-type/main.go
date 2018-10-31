package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Jello world!")
	}
	greeting()
	fmt.Printf("%T\n", greeting)
}
