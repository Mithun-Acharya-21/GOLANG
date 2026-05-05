package main

import (
	"fmt"
	"strconv"
)

func convertString() {
	var i int = 21
	fmt.Printf("String Conversion - Original int: %v, Type: %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("String Conversion - Converted string: %v, Type: %T\n", j, j)

	// Convert string back to int
	converted, err := strconv.Atoi(j)
	if err == nil {
		fmt.Printf("String Conversion - Converted back to int: %v, Type: %T\n", converted, converted)
	}
}