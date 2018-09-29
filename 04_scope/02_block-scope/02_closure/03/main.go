package main

import (
	"fmt"
)

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// closure helps us limit the scope of variable sused by multible functions without closure, for two or mor funcs to have access to the same variable, that variable wold need to be package scope
