package main

import (
	"fmt"
)

/*
func main() {
	var num int = 32767
	num = num + 1
	fmt.Print(num)

	var floatNum float64 = 3.14
	floatNum = floatNum + 0.001
	fmt.Print(floatNum)

	var floatNum32 float32 = 10.1
	var intnum32 int32 = 2
	var result float32 = floatNum32 + float32(intnum32)
	fmt.Print(result)

	var intNum1 int = 3
	var inNum2 int = 2
	fmt.Println(intNum1 / inNum2)
	fmt.Println(intNum1 % inNum2)

	var mySting string = "Hello, \nWorld!"
	fmt.Println(mySting)

	fmt.Println(utf8.RuneCountInString(mySting))

	var myRune rune = 'A'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3)

	const myConst string = "This is a constant string."
	fmt.Println(myConst)
}
*/

// ----------------------------------------------------
/*
func main() {
	printMe()
	adding, subtracting := add_sub(10, 5)

	fmt.Println(adding)
	fmt.Println(subtracting)
	fmt.Printf("Adding: %d, Subtracting: %d\n", adding, subtracting)

	var numerator int = 11
	var denominator int = 0
	var results, remainder, err = intDivison(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Results: %d, Remainder: %d\n", results, remainder)

}

func printMe() {
	fmt.Println("Hello, World!")
}

func add_sub(a int, b int) (int, int) {
	return a + b, a - b
}

func intDivison(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
*/
/// ----------------------------------------------------

/*
func main() {
	var intArr [3]int32
	intArr[1] = 10
	fmt.Println(intArr)

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	/// Or intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// Slices - dynamic arrays
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	// Maps - key-value pairs
	fmt.Printf("Starting maps\n")
	var myMap2 = map[string]uint8{"Adam": 30, "Eve": 28}
	fmt.Println(myMap2["Adam"])
	fmt.Print(myMap2["Eve"])

	// For loops with maps
	fmt.Printf("Starting for loops\n")

	// Looping through key and values - more of a for-each loops - each item
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	// Looping through a list - more of a traditional for loop - index based looping
	for i, v := range intArr2 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
*/
// ----------------------------------------------------

/*
func main() {
	var myString = []rune("resume")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"Go", " is ", "awesome"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
*/
//----------------------------------------------------
// Structs - custom data types
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Ayden"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
}
