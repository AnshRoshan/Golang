package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	var ptr *int
	fmt.Println("Value of pointer ", ptr)
	mynumber := 45
	ptr = &mynumber
	fmt.Println("Value is", ptr)
	fmt.Println("Value is", *ptr)
	*ptr = *ptr + 10
	fmt.Println("Value is", mynumber)

}
