package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")
	var list [4]string

	list[0] = "Blue"
	list[1] = "Pink"
	list[3] = "Green"
	list[3] = "Red"
	//empty element is demoted by space
	fmt.Println(list)

	var name = [3]string{"ansh", "anuj", "alpha"}
	fmt.Println(name)
}
