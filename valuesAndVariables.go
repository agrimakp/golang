package main

import "fmt"

func main() {
	var stationName string
	var nextTrainTime int8
	var isUptownTrain bool

	stationName = "Union Square"
	nextTrainTime = 12
	isUptownTrain = false

	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	stationName = "Grand Central"
	nextTrainTime = 3
	isUptownTrain = true

	fmt.Println("")
	fmt.Println("Current station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	//Literals

	fmt.Println(2235 * 1231)

	//Constants

	const earthsGravity = 9.80665
	// Here's where we print out the gravity:
	fmt.Println(earthsGravity)

	var jellybeanCounter int8
	fmt.Println(jellybeanCounter)

	var favoriteSnack string
	// Assign a value to
	// favoriteSnack
	favoriteSnack = "puff"

	var description string
	description = "My favorite snack is " + favoriteSnack
	fmt.Println(description)

	//Infering type
	// Define daysOnVacation using := below:
	daysOnVacation := 3

	// Define hoursInDay using var and = below:
	var hoursInDay = 24
	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")

	//
	coolSneakers := 65.99
	niceNecklace := 45.50

	// Define taxCalculation here
	var taxCalculation float64
	// Add coolSneakers to taxCalculation
	taxCalculation += coolSneakers
	// Add niceNecklace to taxCalculation
	taxCalculation += niceNecklace
	// Compute the NYC sales tax
	// 8.875% of the purchase here:
	taxCalculation *= .08875
	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)

	//MULTIPLE VARIABLE DECLARATION

	// Define magicNum and powerLevel below:

	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001
	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
}
