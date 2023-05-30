package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n') // if everything is ok, error will be nil
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)

}
