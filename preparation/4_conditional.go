package main

import "fmt"

func print_hello() {
	fmt.Printf("Hello World \n")
}

func main() {

	// defer
	defer print_hello()

	// if, else if, else
	age := 19

	if age < 18 {
		fmt.Printf("Not enough\n")
	} else if age > 60 {
		fmt.Printf("Too old\n")
	} else {
		fmt.Printf("Welcome\n")
	}

	// switch
	grade := "F"
	switch grade {
	case "A":
		fmt.Printf("Good job\n")
	case "B", "C":
		fmt.Printf("Nice work\n")
		fmt.Printf("Not bad\n")
	case "F":
		fmt.Printf("Please see the teacher\n")
		fallthrough
	default:
		fmt.Printf("You have to try again\n")
	}

}
