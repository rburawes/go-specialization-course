/**
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(arr []int) {
	arrayLength := len(arr)
	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j)
			}
		}
	}
}

func swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main() {
	fmt.Println("Please enter up to 10 numbers separated by ',': ")
	scanner := bufio.NewScanner(os.Stdin)
	sli := make([]int, 0, 10)

	if scanner.Scan() {
		inputNumbers := scanner.Text()
		numbers := strings.Split(inputNumbers, ",")
		if len(numbers) <= 10 {
			for _, number := range numbers {

				intNum, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("an error has occurred ", err)
					continue
				}
				sli = append(sli, intNum)

			}
			BubbleSort(sli)
			fmt.Println("The sorted array is: ", sli)
		} else {
			fmt.Println("For now input 10 numbers only! Please, try again.")
		}

	}
}
