package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers in GoLang")

	// var ptr *int // ptr is a pointer to an integer and * is used to declare a pointer
	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23

	var ptr = &myNumber // & is used to get the address of a variable, reference operator
	fmt.Println("Value of myNumber is: ", myNumber)
	fmt.Println("Address of myNumber is: ", &myNumber)
	fmt.Println("Value of ptr is: ", ptr)
	fmt.Println("Value of *ptr is: ", *ptr) // * is used to dereference a pointer

	*ptr = *ptr + 2
	fmt.Println("New value of myNumber is: ", myNumber)

}
