package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name  string
	food  string
	move  string
	sound string
}

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

func get_animal_requests(animal Animal, animal_request string) {
	animal_requests_mapper := map[string]func(){
		"eat":   animal.Eat,
		"move":  animal.Move,
		"speak": animal.Speak,
	}
	if request_func, found := animal_requests_mapper[animal_request]; !found {
		panic("Sorry the animal's request not implemented yet")
	} else {
		request_func()
	}

}

func read_values() []string {
	reader := bufio.NewReader(os.Stdin)

	slice := make([]string, 0, 2)

	fmt.Printf("\n>")
	if values, _, err := reader.ReadLine(); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		values_arr := strings.Split(string(values), " ")
		if len(values_arr) != 2 {
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

func main() {
	fmt.Println("Please insert sequentially animal's name : (cow | bird | snake) and the request's (eat | move | speak) afterwards")
	fmt.Println("Values must inputted separated by SPACE, e.g : cow eat")
	fmt.Println("do CTRL + C to exis the prompt!")
	for {
		// read the vals from inputted
		inputted_vals := read_values()
		animal_name, animal_request := inputted_vals[0], inputted_vals[1]
		animals_struct := get_animals(animal_name)
		get_animal_requests(animals_struct, animal_request)
	}

}
