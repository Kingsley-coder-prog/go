package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	calculateFutureValues(investmentAmount, expectedReturnRate, years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Fututre Value: %.0f\n", futureValue)
	formattedFRV := fmt.Sprintf("Fututre Value (adsjusted for Inflation): %.0f\n", futureRealValue)

	// fmt.Println("Future Value",futureValue)
	// fmt.Printf("Fututre Value %.0f\nFututre Value (adsjusted for Inflation): %.0f", futureValue, futureRealValue)
	// fmt.Print("Fututre Value (adsjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}
