package main

import "fmt"

func main() {
	for k := 0; k < 100; k++ {
		fmt.Printf("%x \t %X", k, k)
	}
}
