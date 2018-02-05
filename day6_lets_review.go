package main

import (
	"fmt"
	"strings"
)

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)

	for i := 0; i < testCases; i++ {
		var str, evenChars, oddChars string
		fmt.Scanf("%v", &str)

		strArr := strings.Split(str, "")
		for i, alphabet := range strArr {
			if i%2 == 0 {
				evenChars += alphabet
			} else {
				oddChars += alphabet
			}
		}

		fmt.Printf("%v %v\n", evenChars, oddChars)
	}
}
