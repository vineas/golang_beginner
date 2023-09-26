package main

import (
	"errors"
	"fmt"
)

func main() {
	var variableNameOne string = "Hello world"
	variableNameTwo := "Haii"
	fmt.Println(variableNameOne)
	fmt.Println(variableNameTwo)

	// boolean
	boolVar := false
	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	// integer
	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	// float
	floatVar := float32(2.9)
	floatVar2 := float64(2.9)
	fmt.Printf("Type: %T Value: %v\n", floatVar, floatVar)
	fmt.Printf("Type: %T Value: %v\n", floatVar2, floatVar2)

	// bytes
	bytesVar := byte(25)
	fmt.Printf("Type: %T Value: %v\n", bytesVar, bytesVar)

	bytesVar2 := []byte("Hello World")
	fmt.Printf("Type: %T Value: %v\n", bytesVar2, bytesVar2)

	// rune
	runeVar := '😏'
	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	// complex
	complexVar := -7 + 3i
	fmt.Printf("Type: %T Value: %v\n", complexVar, complexVar)

	// error
	errVar := errors.New("error detected")
	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)

	// interface
	var myInterfaceVar interface{}

	myInterfaceVar = 5
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)

	myInterfaceVar = "Hello World"
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)

}

type MethodList interface {
	MyFunction()
	MyFunction2(int) int
}
