package main

import (
	"bufio"
	"fmt"
	"os"
)

func hello() {
	reader := bufio.NewReader(os.Stdin)
	if text, err := reader.ReadString('\n'); err != nil {
		fmt.Println("Hello, World.")
		fmt.Println(text)
	}
}
