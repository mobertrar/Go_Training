package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Rob"]) //44
}

func changeMe(z map[string]int) {
	z["Rob"] = 44
}
