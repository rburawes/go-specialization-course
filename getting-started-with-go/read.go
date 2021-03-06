/**
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
Submit your source code for the program, “read.go”.
*/

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
