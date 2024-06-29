package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Golang")
	var list = []string{"pea", "beans"}
	fmt.Printf("Type of list %T\n", list)

	var listarr = [4]string{"ansh", "alpha"}
	fmt.Printf("Type of list %T\n", listarr)

	list = append(list, "Mango", "Banana")
	fmt.Println(list)

	// [inclusive index : exclusive index] for slicing
	list = append(list[1:3])
	fmt.Println(list)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 564
	highScore[2] = 478
	highScore[3] = 912
	//panic: runtime error: index out of range [4] with length 4
	// highScore[4] = 970

	//go behave differently and reallocate the memory
	highScore = append(highScore, 1000, 1024)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var books = []string{"react", "python", "swift", "ruby"}
	fmt.Println(books)
}
