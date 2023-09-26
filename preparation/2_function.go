package main

import "fmt"

func greet(name string) (text string) {
	text = "Halo" + name
	return
}

func add(x, y int) int {
	return x + y
}

func addSub(x, y int) (addRes int, subRes int) {
	// addRes = x + y
	// subRes = x - y
	return x + y, x - y
}

func main() {
	fmt.Println(addSub(2, 4))

}
