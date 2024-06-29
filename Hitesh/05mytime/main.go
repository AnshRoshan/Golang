package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Studying time module")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))
	createDate := time.Date(2000, time.April, 8, 11, 34, 56, 89, time.Local)
	fmt.Println(createDate)
}
