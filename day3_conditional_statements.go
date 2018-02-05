package main

import "fmt"

func conditionalStatements() {
	var n int
	fmt.Scanf("%d\n", &n)

	if n%2 != 0 {
		fmt.Println("Weird")
	} else {
		if 2 < n && n < 5 {
			fmt.Println("Not Weird")
		} else if 6 < n && n < 20 {
			fmt.Println("Weird")
		} else if n > 20 {
			fmt.Println("Not Weird")
		}
	}

}
