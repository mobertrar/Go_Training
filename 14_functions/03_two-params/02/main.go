package main

import "fmt"

func main() {
	fan("front room", "back room")
}

func fan(front, back string) {
	fmt.Println(front, back)
}
