package main

import "fmt"

// Variadic functions
func sumValues(values ...int) int { // A function that receives an indefinite number of parameters
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

// Named returns
func getValues(value int) (double, triple, quad int) { // Defines variables to be returned
	double = value * 2
	triple = value * 3
	quad = value * 4
	return
}

func main() {
	// Variadic functions testing
	fmt.Println(sumValues(2, 2, 3))

	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sumValues(values...))

	// Named return functions testing
	fmt.Println(getValues(2))
}
