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

// var cow, bird, snake Animal = Animal{"cow", "grass", "walk", "moo"},
// 	Animal{"bird", "worms", "fly", "peep"},
// 	Animal{"snake", "mice", "slither", "hsss"}

var ANIMALS_MAPPER = map[string]Animal{
	"cow":   {"cow", "grass", "walk", "moo"},
	"bird":  {"bird", "worms", "fly", "pheep"},
	"snake": {"snake", "mice", "slither", "shsss"},
}

func (animal *Animal) Eat() {
	name, food := animal.name, animal.food
	fmt.Printf("The animal %v eat is %v\n", name, food)
}

func (animal *Animal) Move() {
	name, move := animal.name, animal.move
	fmt.Printf("The animal %v move is %v\n", name, move)
}

func (animal *Animal) Speak() {
	name, sound := animal.name, animal.sound
	fmt.Printf("The animal %v sound like %v\n", name, sound)
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
		panic("Sorry the animal's request not implemented yet")
	}
}

func insert_new_animal() {
	var name, eat, move, speak string

	fmt.Println("Please input your new animal separate by space and ordering by name, eat, move, speak, e.g: tiger meat walk roar")
	fmt.Scan(&name, &eat, &move, &speak)

	ANIMALS_MAPPER[name] = Animal{name, eat, move, speak}
	fmt.Printf("Successfully create new animal, now you can type e.g: %v eat \n", name)
}

func read_values() []string {
	reader := bufio.NewReader(os.Stdin)

	slice := make([]string, 0, 2)

	fmt.Printf("\n>")
	if values, _, err := reader.ReadLine(); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		values_arr := strings.Split(string(values), " ")
		_, animal_found := ANIMALS_MAPPER[values_arr[0]]
		if len(values_arr) != 2 && animal_found {
			panic("The inputted values must be filled in: {animal's name} {request's name}")
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
		panic("The animal's name you inputted not registered")
	}
	return val
}

func get_animal_list(animals []string) string {
	for k := range ANIMALS_MAPPER {
		animals = append(animals, k)
	}
	return strings.Join(animals, " | ")
}

func main() {
	animals := make([]string, 0, len(ANIMALS_MAPPER))

	for {
		animal_list := get_animal_list(animals)
		fmt.Println("======================")
		fmt.Printf("Please insert sequentially animal's name : (%v) and the request's (eat | move | speak) afterwards", animal_list)
		fmt.Println("Values must inputted separated by SPACE, e.g : cow eat")
		fmt.Println("OR you likely to create your own animale by type 'new animal'!")
		fmt.Println("do CTRL + C to exis the prompt!")
		// read the vals from inputted
		inputted_vals := read_values()
		animal_name, animal_request := inputted_vals[0], inputted_vals[1]
		if animal_name == "new" {
			insert_new_animal()
			continue
		}
		animals_struct := get_animals(animal_name)
		get_animal_requests(&animals_struct, animal_request)
	}

}
