package main

import (
	"fmt"
)

func main() {
	var name string
	var number float64

	fmt.Print("Enter a name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	fmt.Printf("Hi %v, your number is %v\n", name, number) // need '\n' so % doesn't print
}
