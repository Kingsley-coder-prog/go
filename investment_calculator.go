package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Fututre Value: %.0f\n", futureValue)
	formattedFRV := fmt.Sprintf("Fututre Value (adsjusted for Inflation): %.0f\n", futureRealValue)

	// fmt.Println("Future Value",futureValue)
	// fmt.Printf("Fututre Value %.0f\nFututre Value (adsjusted for Inflation): %.0f", futureValue, futureRealValue)
	// fmt.Print("Fututre Value (adsjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}
