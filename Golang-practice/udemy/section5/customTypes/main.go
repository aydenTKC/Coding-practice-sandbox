package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(string(text))
}
func main() {
	var name str = "Ayden"

	name.log()

}
