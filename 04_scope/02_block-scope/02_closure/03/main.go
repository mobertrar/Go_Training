package main

import (
	"fmt"
)

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	} //the purple curlys are the inner scope
	fmt.Println(increment())
	fmt.Println(increment())
} //the yellow curlys are the outer scope

// closure helps us limit the scope of variable sused by multible functions without closure, for two or mor funcs to have access to the same variable, that variable wold need to be package scope

//anonymous function - a function without a name
// example  func ()

//function expression - assigning a functions to a variable
// example increment := func ()
