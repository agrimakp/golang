package main

import (
	"fmt"
)

func main() {
	name1 := "Shidil"
	name2 := "Shidil Eringa"
	var name string
	//1
	fmt.Println("Hello World")
	//2
	fmt.Print("1: Hello ", name1)
	fmt.Println("")

	fmt.Printf("2: Hello %s", name2)
	fmt.Println("")

	// 3
	fmt.Println("Your Name ?")
	fmt.Scan(&name)
	fmt.Printf("Name: %s", name)
	fmt.Println("")

}
