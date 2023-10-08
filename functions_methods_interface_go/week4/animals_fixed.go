package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal_interface interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	name  string
	food  string
	move  string
	sound string
}

var cow, bird, snake Animal = Animal{"cow", "grass", "walk", "moo"},
	Animal{"bird", "worms", "fly", "peep"},
	Animal{"snake", "mice", "slither", "hsss"}

var ANIMALS_MAPPER = map[string]Animal{}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.move)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.sound)
}

func get_animal_requests(animal *Animal, animal_request string) {
	map_animal_requests(animal, animal_request)
}

func map_animal_requests(animal animal_interface, animal_request string) {
	switch animal_request {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Sorry the animal's request not implemented yet")
	}
}

func insert_new_animal(name string, animal_type string) {
	msg := "Created it! now you can use command e.g: query %v eat\n"
	_, found := ANIMALS_MAPPER[name]
	if !found {
		switch animal_type {
		case "cow":
			ANIMALS_MAPPER[name] = cow
			fmt.Printf(msg, name)
		case "bird":
			ANIMALS_MAPPER[name] = bird
			fmt.Printf(msg, name)
		case "snake":
			ANIMALS_MAPPER[name] = snake
			fmt.Printf(msg, name)
		default:
			fmt.Println("animals can be registered are cow/bird/snake")
		}
	} else {
		fmt.Printf("The animal with name %v already exists it registered as %v\n", name, animal_type)
	}

}

func read_values() []string {
	reader := bufio.NewReader(os.Stdin)

	slice := make([]string, 0, 3)

	fmt.Printf("\n>")
	if values, _, err := reader.ReadLine(); err != nil {
		fmt.Printf("Error when inputting command: %v", err)
	} else {
		values_arr := strings.Split(string(values), " ")
		if len(values_arr) != 3 {
			fmt.Println("Wrong command input please read the instructions first")
		}

		for _, input := range values_arr {
			slice = append(slice, string(input))
		}
	}
	return slice
}

func get_animals(name string) Animal {
	val, animal_found := ANIMALS_MAPPER[name]
	if !animal_found {
		fmt.Printf("The animal's name: %v you inputted not registered", name)
	}
	return val
}

func main() {
	for {
		fmt.Println("======================")
		fmt.Println("Register new animal using this command 'newanimal <animal-name> cow/bird/snake' e.g: newanimal bob cow")
		fmt.Println("OR just query to get existing animal using this command 'query <animal-name> eat/move/speak' e.g: query bob eat")
		fmt.Println("do CTRL + C to exit the prompt!")
		// read the vals from inputted
		inputted_vals := read_values()
		command := inputted_vals[0]
		switch command {
		case "newanimal":
			animal_name, animal_type := inputted_vals[1], inputted_vals[2]
			insert_new_animal(animal_name, animal_type)
		case "query":
			animal_name, animal_request := inputted_vals[1], inputted_vals[2]
			animals_struct := get_animals(animal_name)
			get_animal_requests(&animals_struct, animal_request)
		default:
			fmt.Println("The command only either 'newanimal' or 'query")
		}
	}

}
