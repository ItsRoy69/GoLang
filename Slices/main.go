package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to a class on Slices in GoLang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of fruit list is: ", fruitList)     // [Apple Tomato Peach]
	fmt.Println("Fruit list length is: ", len(fruitList)) // 3 (length of the slice)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit list after appending is: ", fruitList) // [Apple Tomato Peach Mango Banana]

	fruitList = append(fruitList[1:3])                      // [Tomato Peach] as 1 is inclusive and 3 is exclusive
	fmt.Println("Fruit list after slicing is: ", fruitList) // [Tomato Peach]

	highScores := make([]int, 4) // make is used to create a slice of length 4
	highScores[0] = 234          // [234 0 0 0]
	highScores[1] = 945          // [234 945 0 0]
	highScores[2] = 465          // [234 945 465 0]
	highScores[3] = 867          // [234 945 465 867]
	// highScores[4] = 777          // OUT OF BOUNDS ERROR as length of slice is 4

	highScores = append(highScores, 555, 666, 777) // [234 945 465 867 555 666 777] as length of slice is now 7 after appending
	fmt.Println("High scores are: ", highScores)   // [234 945 465 867 555 666 777]

	sort.Ints(highScores) // [234 465 555 666 777 867 945] sorts the slice
	fmt.Println("Sorted high scores are: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // true as the slice is sorted

	//how to remove a value from a slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Course list is: ", courses) // [reactjs javascript swift python ruby]
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // [reactjs javascript python ruby] as swift is removed from index 2 (0 based indexing) and ... is used to append the slice, colon index is inclusive and colon index+1 is exclusive
	fmt.Println("Course list after removing index 2 is: ", courses)

}
