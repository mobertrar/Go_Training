//Package stringutil contains utility functions for working with strings.
package stringutil

//Reverse returs its argument string reversed run-wise left to right
func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build

				go build revers.go reverseTwo.go
				won't produce an output file

go install

				will place the package inside the pkg directory of the workspace
*/
