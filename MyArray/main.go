package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Arrays in GoLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)             // [Apple Tomato  Peach]
	fmt.Println("Fruit list length is: ", len(fruitList)) // 4 (length of the array)

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list is: ", vegList) // [Potato Beans Onion]

}
