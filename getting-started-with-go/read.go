package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstname, lastname string
}

/**
Please provide the name of the file to be scanned.
data.txt
The names from the file:
Raul Banchero
Alfred Vargas
James Yap
LA Tenorio
*/
func main() {
	fmt.Println("Please provide the name of the file to be scanned.")
	var filename string
	_, err := fmt.Scanf("%s", &filename)
	if err != nil {
		log.Fatal("Unable to read the provided file ", err)
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error opening file. err is ", err)
		return
	}
	persons := []Person{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedBytes := scanner.Bytes()
		byteArr := make([]byte, 20)
		copy(byteArr, scannedBytes)
		for index, b := range byteArr {
			if rune(b) == '\n' {
				byteArr = byteArr[:index]
				break
			}
		}
		names := strings.Split(string(byteArr), " ")

		if len(names) >= 2 {
			persons = append(persons, Person{
				firstname: names[0],
				lastname:  names[1],
			})
		}
	}
	if err := file.Close(); err != nil {
		fmt.Println("An error has occurred while closing the file ", err)
	}

	fmt.Println("The names from the file: ")
	for _, person := range persons {
		fmt.Println(person.firstname, person.lastname)
	}
}
