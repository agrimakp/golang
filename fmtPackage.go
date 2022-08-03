package main

import "fmt"

func main() {
	fmt.Println("There's default space", "between the string arguments.")
	fmt.Print("Print", "is", "different")
	fmt.Print("See?")

	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)
	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)

	//Sprint
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	meditation := fmt.Sprintln(step1, step2)
	fmt.Print(meditation)

	// ----------------------------------------
	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)
	// -----------------------------scan()
	fmt.Println("What would you like for lunch?")

	// Add your code below:
	var food string
	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)
}
