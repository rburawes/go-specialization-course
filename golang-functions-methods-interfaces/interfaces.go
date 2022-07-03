/**
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	| Food eaten | Locomotion method | Spoken sound
cow		  grass 	   walk		  		   moo
bird 	  worms 	   fly 		  		   peep
snake 	  mice 		   slither    		   hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.
Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.
Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

Submit your Go program source code.
*/
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
