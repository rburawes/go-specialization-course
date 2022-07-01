package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Please enter any string: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Input string not found")
	}
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
