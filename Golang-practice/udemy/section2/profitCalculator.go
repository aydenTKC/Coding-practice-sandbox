package main

import (
	"fmt"
)

func main() {
	var revenue int
	var expense int
	var taxRate float64

	fmt.Print("What is your monthly revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is your monthly expense: ")
	fmt.Scan(&expense)

	fmt.Print("What is your monthly tax rate: ")
	fmt.Scan(&taxRate)

	// EBT = Earning before tax
	ebt := revenue - expense
	profit := float64(revenue) / float64(expense)

	// EAT = Earning after tax
	eat := revenue - (taxRate * revenue)

}
