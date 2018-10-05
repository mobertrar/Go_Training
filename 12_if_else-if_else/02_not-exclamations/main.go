package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This is did not run")
	}
	if !false {
		fmt.Println("This ran")
	}
}
