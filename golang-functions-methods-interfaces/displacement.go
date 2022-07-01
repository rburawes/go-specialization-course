package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, v, s float64
	var t int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the acceleration value:")
	if !scanner.Scan() {
		return
	}
	a, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter initial velocity value:")
	if !scanner.Scan() {
		return
	}
	v, err = strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter initial displacement:")
	if !scanner.Scan() {
		return
	}
	s, err = strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter time value in seconds:")
	if !scanner.Scan() {
		return
	}
	t, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	fn := GenDisplaceFn(a, v, s)
	fmt.Println("Displacement after ", t, "seconds is: ", fn(t))
}

func GenDisplaceFn(acceleration, velocity, initDisplacemnt float64) func(int) float64 {
	fn := func(t int) float64 {
		return (0.5 * acceleration * math.Pow(float64(t), 2.0)) + initDisplacemnt + (velocity * float64(t))
	}
	return fn
}
