package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	vals := read_values()

	fn := gen_displace_fn(vals...)

	result := func(t float64, n func(float64) float64) {
		fmt.Printf("\nDisplacement after %v seconds is %v.", []any{t, fn(t)}...)
	}

	// 3 seconds
	result(3.0, fn)
	// 5 seconds
	result(5.0, fn)

}

func read_values() []float64 {
	reader := bufio.NewReader(os.Stdin)

	slice := make([]float64, 0, 3)

	fmt.Println("Please insert sequentially from : Acceleration, Velocity and Displacement values (Must INTEGER/FLOAT)")
	fmt.Println("Values must inputted separated by SPACE, e.g : 1.2 9.0 8.2")
	if values, _, err := reader.ReadLine(); err != nil {
		fmt.Println("Error")
	} else {
		values_arr := strings.Split(string(values), " ")
		if len(values_arr) != 3 {
			panic("The inputted values must be filled in: Acceleration, Velocity and Displacement")
		}

		for _, num := range values_arr {
			if num_float, err := strconv.ParseFloat(num, 64); err != nil {
				panic(fmt.Sprintf("Number that you inputted must be INTEGER/FLOAT: %v", num))
			} else {
				slice = append(slice, num_float)
			}

		}
	}
	return slice
}

func gen_displace_fn(vals ...float64) func(float64) float64 {
	a, v0, s0 := vals[0], vals[1], vals[2]

	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}
