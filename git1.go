package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	var name string
	fmt.Println("Input username")
	fmt.Scan(&name)

	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println()

	fmt.Println("Name:")

}
