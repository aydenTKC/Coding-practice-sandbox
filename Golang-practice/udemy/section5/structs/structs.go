package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name.")
	userlastName := getUserData("Please enter your last name.")
	userbirthdate := getUserData("Please enter your birth date.")

	var appUser *user.User

	appUser, err := user.NewUser(userfirstName, userlastName, userbirthdate)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gather data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promtText string) string {
	fmt.Println(promtText)
	var value string
	// Scanln = with the ln takes empty inputs, so enter/next line
	fmt.Scanln(&value)
	return value
}
