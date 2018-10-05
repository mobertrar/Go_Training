package main

import "fmt"

func main() {
	counter := 0
	for v := 0; v < 10; v++ {
		if v%3 == 0 {
			counter += v

		} else if v%5 == 0 {
			counter += v

		}

	}
	fmt.Println(counter)
}
