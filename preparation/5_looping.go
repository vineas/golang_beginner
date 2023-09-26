package main

import "fmt"

func main() {

	// for
	for i := 0; i < 5; i++ {
		fmt.Printf("Looping %d\n", i)
	}

	// while
	j := 0
	for j < 5 {
		fmt.Printf("Looping %d\n", j)
		j += 1
	}

	// infinte loop, break, and continue
	i := 0
	for {
		i += 1
		if i == 2 {
			continue
		}
		fmt.Printf("Looping %d\n", i)
		if i == 5 {
			break
		}
	}

	// range loop iterator
	numbers := []int64{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("I got %d on index %d\n", numbers[i], i)
	}

}
