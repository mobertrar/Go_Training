package main

import "fmt"

func main() {
	fmt.Println(greet("Josh", "Howard"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
