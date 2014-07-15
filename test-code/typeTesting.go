package main

import (
	"fmt"
)

type floatSlice []float64

func main() {
	A := floatSlice{1, 2, 3, 4, 5}

	var B floatSlice
	B = floatSlice{98, 93, 77, 82, 83}

	fmt.Println(A, B)
}
