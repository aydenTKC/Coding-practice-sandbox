package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	const inflationRate = 2.5

	fmt.Print("Give me investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Give me years of investing: ")
	fmt.Scan(&years)

	fmt.Print("What is the return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
