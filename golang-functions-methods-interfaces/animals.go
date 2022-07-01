package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	diet     string
	movement string
	sound    string
}

func (a Animal) Eat() {
	fmt.Println(a.diet)
}

func (a Animal) Move() {
	fmt.Println(a.movement)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func initAnimals() map[string]*Animal {
	cow := &Animal{
		diet:     "grass",
		movement: "walk",
		sound:    "moo",
	}
	bird := &Animal{
		diet:     "worms",
		movement: "fly",
		sound:    "peep",
	}
	snake := &Animal{
		diet:     "mice",
		movement: "slither",
		sound:    "hsss",
	}
	m := make(map[string]*Animal)
	m["cow"] = cow
	m["bird"] = bird
	m["snake"] = snake
	return m
}

func main() {
	objectMap := initAnimals()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pick an animal [cow,bird,snake]: ")
		scanner.Scan()
		object := strings.ToLower(scanner.Text())
		if _, ok := objectMap[object]; !ok {
			fmt.Println("Sorry, animal is not in the options.")
			break
		}
		fmt.Println("Select option to know more about the animal [eat, move, speak]")
		scanner.Scan()
		about := strings.ToLower(scanner.Text())

		switch about {
		case "eat":
			objectMap[object].Eat()
		case "move":
			objectMap[object].Move()
		case "speak":
			objectMap[object].Speak()
		default:
			fmt.Println("Sorry, the option you've selected is not in the list.")
		}
	}
}
