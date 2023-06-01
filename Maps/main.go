package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Maps in GoLang")

	languages := make(map[string]string) // make is used to create a map
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("Languages map is: ", languages) // map[GO:Golang JS:Javascript PY:Python RB:Ruby]
	fmt.Println("Length of languages map is: ", len(languages))

	// delete(languages, "RB") // delete is used to delete a key value pair from a map
	// fmt.Println("Languages map after deleting RB is: ", languages) // map[GO:Golang JS:Javascript PY:Python]

	// for key, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", key, value)
	// }

}
