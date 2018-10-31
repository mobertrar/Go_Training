package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Jello World!")
	}
	greeting()
}
