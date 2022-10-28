package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Type userName")
	fmt.Scan(&name)
	fmt.Println("Your github URL : ")
	fmt.Printf("https://github.com/%s", name)
	fmt.Println("")

}
