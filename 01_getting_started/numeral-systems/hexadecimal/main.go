package main

import "fmt"

func main() {
	for k := 0; k < 10; k++ {
		fmt.Printf("%x \t %X", k, k)
	}
	//for l := 0; l < 10; l++ {
		//fmt.Printf("%d \n %d", l, l)
	//}
}
