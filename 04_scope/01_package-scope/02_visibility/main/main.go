package main

import (
	"fmt"

	"github.com/GoesToEleven/Go_Training/scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
