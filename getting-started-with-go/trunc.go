package main

import "fmt"

func main() {
	var fNum float64
	fmt.Print("Please enter floating point number: ")
	fmt.Scan(&fNum)
	iNum := int(fNum)
	fmt.Println("The integer value of the floating point number is: ", iNum)
}
