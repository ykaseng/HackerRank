package main

import (
	"fmt"
	"math"
)

func operators() {
	var (
		mealCost   float64
		tipPercent int
		taxPercent int
	)

	fmt.Scanf("%f\n", &mealCost)
	fmt.Scanf("%d\n", &tipPercent)
	fmt.Scanf("%d\n", &taxPercent)

	fmt.Printf("The total meal cost is %v dollars.", round(mealCost+(float64(tipPercent)/100*mealCost)+(float64(taxPercent)/100*mealCost)))
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
