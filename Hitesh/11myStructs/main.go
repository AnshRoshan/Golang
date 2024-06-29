package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	// no inheritance
	ansh := User{"Ansh", "ansh@dev.com", true, 22}
	fmt.Println(ansh)
	fmt.Println(ansh.email)
	fmt.Printf("Ansh Details are: %+v\n", ansh)

}

type User struct {
	name   string
	email  string
	status bool
	age    int
}
