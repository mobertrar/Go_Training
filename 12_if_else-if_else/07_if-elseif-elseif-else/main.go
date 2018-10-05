package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print statment")
	} else if !false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahahahaha print statement")
	} else {
		fmt.Println("third print statement")
	}
}
