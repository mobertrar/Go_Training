package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}

//Some operating systems (Windows) might not print characters where i<256
//if you have this issue, you can use htis code:
//fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

//UTF-8 is the text coding scheme used by Go.

//UTF-8 works with 1 - 4 bytes.
