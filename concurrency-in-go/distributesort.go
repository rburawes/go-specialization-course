/**
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortNumbers(arr []int, goroutine int, output chan []int) {
	fmt.Println("Before sort operation: ", arr, " [goroutine]: ", goroutine)
	sort.Ints(arr)
	fmt.Println("Before sort operation: ", arr, " [goroutine]: ", goroutine)
	output <- arr
}

func merge(arr1, arr2 []int) []int {
	var arr []int
	n1, n2 := len(arr1), len(arr2)
	var i, j int
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			arr = append(arr, arr2[j])
			j++
		} else {
			//if both equal
			arr = append(arr, arr1[i])
			arr = append(arr, arr2[j])
			i, j = i+1, j+1
		}
	}
	for i < n1 {
		arr = append(arr, arr1[i])
		i++
	}
	for j < n2 {
		arr = append(arr, arr2[j])
		j++
	}
	return arr
}

func execute(arr [][]int) []int {
	output := make(chan []int, 4)
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sortNumbers(arr[i-1], i, output)
		}(i)
	}
	wg.Wait()
	close(output)
	wg.Add(2)
	result := make(chan []int, 2)
	go func(result chan []int) {
		defer wg.Done()
		arrs := make([][]int, 0)
		for i := 0; i < 2; i++ {
			temp := <-output
			arrs = append(arrs, temp)
		}
		temp := merge(arrs[0], arrs[1])
		result <- temp
	}(result)
	go func(result chan []int) {
		defer wg.Done()
		arrs := make([][]int, 0)
		for i := 0; i < 2; i++ {
			temp := <-output
			arrs = append(arrs, temp)
		}
		temp := merge(arrs[0], arrs[1])
		result <- temp
	}(result)
	wg.Wait()
	close(result)
	var mergedArrs [][]int
	for arr := range result {
		mergedArrs = append(mergedArrs, arr)
	}
	final := merge(mergedArrs[0], mergedArrs[1])
	fmt.Println("Merged and sorted: ", final)
	return final
}

func main() {
	fmt.Println("Please enter comma separated numbers. e.g. 1,2,3,5,6,8,1,3,4")
	var inputData string
	fmt.Scan(&inputData)
	elements := strings.Split(inputData, ",")
	var arr []int
	for _, ele := range elements {
		e, err := strconv.Atoi(ele)
		if err != nil {
			fmt.Errorf("ERROR: Cannot convert string to number. %v", err)
			return
		}
		arr = append(arr, e)
	}
	n := len(arr)
	partitionSize := int(math.Ceil(float64(n) / 4.0))
	fmt.Println("Partition size: ", partitionSize)
	arr2 := make([][]int, 4)
	for j := 0; j < 4; j++ {
		arr2[j] = make([]int, 0)
	}
	for i, j := 0, -1; i < n && j < 4; i++ {
		if i%partitionSize == 0 {
			j++
		}
		arr2[j] = append(arr2[j], arr[i])
	}
	fmt.Println("The sorted array is ", execute(arr2))
}
