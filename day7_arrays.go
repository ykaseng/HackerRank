package main

import (
	"fmt"
)

func arrays() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var arraySize int
	fmt.Scanf("%d", &arraySize)

	var ints = make([]int, arraySize)
	for i := (arraySize - 1); i >= 0; i-- {
		fmt.Scanf("%d", &ints[i])
	}

	for _, val := range ints {
		fmt.Printf("%d ", val)
	}
}
