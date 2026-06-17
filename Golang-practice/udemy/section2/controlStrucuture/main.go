package main

import (
	"fmt"
	"log"
	"os"      // lets me write and red from a file
	"strconv" // String conversion
)

const acountBalancFile = "balance.txt"

func getBalanceToFile() float64 {
	data, err := os.ReadFile(acountBalancFile)
	if err != nil {
		log.Fatal(err)
	}

	// Read in the file as a string
	balanceText := string(data)

	// Parsing a string into a float
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		log.Fatal(err)
	}

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(acountBalancFile, []byte(balanceText), 0644)
}

func main() {
	var balance = getBalanceToFile()
	choice := 0

	for choice != 4 {
		fmt.Println("Welcome to the bank")
		fmt.Println("What can I do for you today")
		fmt.Println("Pick a number")
		fmt.Println("1 - Deposit money ")
		fmt.Println("2 - Withdraw money ")
		fmt.Println("3 - Check account")
		fmt.Println("4 - Exit")
		fmt.Println("Choose a number: ")
		fmt.Scan(&choice)

		if choice == 1 {
			var deposit float64
			fmt.Println("Welcome to deposit money")
			fmt.Println("How much would you like to deposit: ")
			fmt.Scan(&deposit)
			total := deposit + balance
			fmt.Println("Total balance:", total)
			writeBalanceToFile(total)
		}

		if choice == 2 {
			var withdraw float64
			fmt.Println("Welcome to withdraw money")
			fmt.Println("How much would you like to withdraw: ")
			fmt.Scan(&withdraw)
			total := balance - withdraw
			fmt.Println("Total balance:", total)
			writeBalanceToFile(total)
		}

		if choice == 3 {
			fmt.Println("Welcome to account")
			fmt.Println("You have", balance, "money in your account.")
		}
		if choice == 4 {
			fmt.Println("GoodBye")
		}
	}
}
