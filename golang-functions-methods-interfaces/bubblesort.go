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
