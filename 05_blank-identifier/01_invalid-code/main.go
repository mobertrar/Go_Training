package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a-", a)
	// b is not used - invalid code
	//code is 'polluted' because certain items are unused
}
