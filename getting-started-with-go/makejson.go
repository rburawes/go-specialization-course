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
