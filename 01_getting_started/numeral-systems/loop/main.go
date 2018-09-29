package main

import "fmt"

func main() {
	for a := 0; a < 20; a++ {
		fmt.Printf("%d \t %b \t %x \t %X", a, a, a, a)
	}
}
