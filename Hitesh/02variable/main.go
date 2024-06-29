package main

import "fmt"

const Keyapi = "adasda7as8s9a8sd8aa"             //Public
const privatekey = "vnfdjibvfvieibvuievbeibd8aa" //Public

func main() {
	var username string = "Ansh Roshan"
	fmt.Println(username)
	fmt.Printf("Variable of type : %T \n ", username)
	var isCorrect bool = true
	fmt.Println(isCorrect)
	fmt.Printf("Variable of type : %T \n ", isCorrect)
	var unsignInt uint8 = 255
	fmt.Println(unsignInt)
	fmt.Printf("Variable of type : %T \n ", unsignInt)
	var signInt int8 = 127
	fmt.Println(signInt)
	fmt.Printf("Variable of type : %T \n ", signInt)
	var floatsmall float32 = 25.7821871229812389123
	fmt.Println(floatsmall)
	fmt.Printf("Variable of type : %T \n ", floatsmall)
	var floatlarge float64 = 25.7821871229812389123
	fmt.Println(floatlarge)
	fmt.Printf("Variable of type : %T \n ", floatlarge)
	var complexsm complex64 = 67 + 98978122324234324i
	fmt.Println(complexsm)
	fmt.Printf("Variable of type : %T \n ", complexsm)
	var complexLarg complex128 = 67 + 98978122324234324i
	fmt.Println(complexLarg)
	fmt.Printf("Variable of type : %T \n ", complexLarg)

	//default value and alias
	var defint int
	fmt.Println(defint)
	fmt.Printf("Variable of type : %T \n ", defint)
	var deffloat float64
	fmt.Println(deffloat)
	fmt.Printf("Variable of type : %T \n ", deffloat)

	//implicit variable declartion
	// it can be used globally
	var name = "Ansh Roshan"
	fmt.Println(name)

	//no var type or walrus type
	//walrus is only avilable in method and not be used global
	num := 89.90
	fmt.Println(num)

	fmt.Println(Keyapi)
	fmt.Printf("Variable of type : %T \n ", Keyapi)
	fmt.Println(privatekey)
	fmt.Printf("Variable of type : %T \n ", privatekey)
	fmt.Printf("Variable of type : %#v \n ", privatekey)
	fmt.Printf("Variable of type : %v \n ", privatekey)

}
