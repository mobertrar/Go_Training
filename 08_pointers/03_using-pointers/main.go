package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) //43

	*b = 42        //b says, "The value at this address, chage it to 42"
	fmt.Println(a) //42

	//this is useful
	//we can pass a memory address instead of a bunch of values(we can pass a reference)
	//this makes our programs more performant
	//we don't have to pass around large abouts of data
	//we only have to pass around addresses
}
