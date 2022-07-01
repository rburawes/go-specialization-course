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
