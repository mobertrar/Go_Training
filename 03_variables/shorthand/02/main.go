package main

import "fmt"

func main() {
	a := 10 //int
	b := "golang" //string
	c := 4.17 //float
	d := true //boolean
	e := "Hello" //string
	f := `Do you like my hat?` //string
	g := 'M' //int

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
