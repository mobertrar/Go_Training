package main

import "fmt"

func main() {
	data := []float64{6, 4, 5, 3, 2}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
