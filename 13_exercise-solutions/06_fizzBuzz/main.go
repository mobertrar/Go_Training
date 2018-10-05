package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%15 == 0 {
			fmt.Println(x, "--FizzBuzz")

		} else if x%3 == 0 {
			fmt.Println(x, "--Fizz")
		} else if x%5 == 0 {
			fmt.Println(x, "--Buzz")
		} else {
			fmt.Println(x)
		}
	}
}
