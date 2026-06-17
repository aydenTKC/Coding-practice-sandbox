package main

import (
	"fmt"
)

func main() {
	revenue, expense, taxRate := input()

	//fmt.Print("What is your monthly revenue: ")
	//fmt.Scan(&revenue)

	//fmt.Print("What is your monthly expense: ")
	//fmt.Scan(&expense)

	//fmt.Print("What is your monthly tax rate: ")
	//fmt.Scan(&taxRate)

	// EBT = Earning before tax
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ration := ebt / profit

	fmt.Println("This is the earning before tax: ", ebt)
	fmt.Println(profit)
	fmt.Println(ration)
}

func input() (float64, float64, float64) {
	var revenue, expense, taxRate float64

	fmt.Print("What is your monthly revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is your monthly expense: ")
	fmt.Scan(&expense)

	fmt.Print("What is your monthly tax rate: ")
	fmt.Scan(&taxRate)

	return revenue, expense, taxRate
}
