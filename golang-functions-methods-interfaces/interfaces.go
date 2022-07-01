package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal2 interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println(fmt.Sprintf("%s eats worm!", b.name))
}

func (b Bird) Move() {
	fmt.Println(fmt.Sprintf("%s flies!", b.name))
}
func (b Bird) Speak() {
	fmt.Println(fmt.Sprintf("%s sounds twit!!", b.name))
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println(fmt.Sprintf("%s eats grass!", c.name))
}

func (c Cow) Move() {
	fmt.Println(fmt.Sprintf("%s walk!", c.name))
}

func (c Cow) Speak() {
	fmt.Println(fmt.Sprintf("%s sounds moo!", c.name))
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println(fmt.Sprintf("%s eats mice or small birds!", s.name))
}

func (s Snake) Move() {
	fmt.Println(fmt.Sprintf("%s moves like slither", s.name))
}

func (s Snake) Speak() {
	fmt.Println(fmt.Sprintf("%s sounds hisss!", s.name))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animals := make(map[string]Animal2)
	for {
		fmt.Println()
		fmt.Println("Welcome to the farm! Here you can inquire or create a new animal using 'newanimal' or 'query' commands.")
		fmt.Println(">>>> Available animal are [cow, bird, snake]...")
		fmt.Println(">>>> Animal info that can be used are [speak, move, eat]")
		fmt.Println(">>>> To add new animal type: <command> <name-name-of-the-animal> <animal-type> e.g. newanimal James cow")
		fmt.Println(">>>> To query an animal: <command> <name-name-of-the-animal> <animal-type> e.g. query James eat")
		scanner.Scan()
		command := scanner.Text()
		commands := strings.Split(command, " ")
		if len(commands) != 3 {
			fmt.Println("Unrecognized command! Please enter 3 string command.")
			break
		}
		switch strings.ToLower(commands[0]) {
		case "newanimal":
			var an Animal2
			name := commands[1]
			animal := commands[2]
			switch strings.ToLower(animal) {
			case "cow":
				an = Cow{name: name}
			case "bird":
				an = Bird{name: name}
			case "snake":
				an = Snake{name: name}
			default:
				fmt.Println("Error: Invalid animal, allowed are [cow, bird, snake].")
			}
			animals[name] = an
			fmt.Printf("\nGreat! %s joined the farm.\n", name)
		case "query":
			name := commands[1]
			action := commands[2]
			var animal Animal2
			animal, ok := animals[name]
			if !ok {
				fmt.Println(fmt.Sprintf("Animal with name %s is not available.", name))
				break
			}
			switch strings.ToLower(action) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Error: Invalid animal, allowed are [cow, bird, snake].")
			}
		default:
			fmt.Println("Error: Invalid first command. Please select [newanimal or query].")
		}
	}
}
