package main

import "fmt"

func main() {
	for g := 60; g < 122; g++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", g, g, g, g)
	}
}
