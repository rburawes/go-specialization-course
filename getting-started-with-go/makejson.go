/**
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
Submit your source code for the program,
“makejson.go”.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	info := make(map[string]string)
	fmt.Println("Please enter your name: ")
	if scanner.Scan() {
		inputName := scanner.Text()
		info["name"] = inputName
	}
	fmt.Println("Please enter your address: ")
	if scanner.Scan() {
		inputAddress := scanner.Text()
		info["address"] = inputAddress
	}
	obj, err := json.Marshal(info)
	if err != nil {
		log.Fatal("Unable to convert map to JSON - ", err)
		return
	}
	fmt.Println("JSON value is: ", string(obj))
}
