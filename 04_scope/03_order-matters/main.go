package main

import "fmt"

func main() {

	//fmt.Println(x)
	fmt.Println(y)
	//x := 42
}

var y = 42

//order matters,because var y was assigned after the fmt.Println it was able to be compiled. X was not compiled because it is within the scope and after the print
