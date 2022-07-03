/**
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

Submit your source code for the program,
“slice.go”.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)
	var numStr string
	for {
		fmt.Println("Please input a number, or x to exit.")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numStr = scanner.Text()
		}
		numStr = strings.ToLower(numStr)
		if strings.Compare(numStr, "x") == 0 {
			break
		}
		in, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("unable to convert the input to string ", err)
			continue
		}
		sli = append(sli, in)
		sort.Ints(sli)
		fmt.Printf("'sli' sorted values: %v \n", sli)
	}
}
