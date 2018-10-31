package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Wello Horld!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
