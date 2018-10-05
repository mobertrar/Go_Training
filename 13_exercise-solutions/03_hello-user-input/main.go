package main

import "fmt"

func main() {
	var username string
	fmt.Print("Please Enter you username: ")
	fmt.Scan(&username)
	fmt.Println("Hello", username)

}
