package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ask := "Welcome to new prg"
	fmt.Println(ask)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")

	//comma ok || comma error  syntax
	// input, err := reader.ReadString('\n')
	// underscore if u dont care for the value
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for giving %v rating.", input)
}

