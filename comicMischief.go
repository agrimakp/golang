package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int16
	var grade float32
	title = "Mr. GoToSleep"
	publisher = "DizzyBooks Publishing Inc."
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	year = 1997
	pageNumber = 14
	grade = 6.5
	fmt.Println(title, "written by", writer, "drawn by", artist, "Published by", publisher, "in", year, "total page number is", pageNumber, "grade =", grade)

	title = "Epic Vol. 1"
	publisher = "DizzyBooks Publishing Inc."
	writer = "Ryan N. Shawn."
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0
	fmt.Println(title, "written by", writer, "drawn by", artist, "Published by", publisher, "in", year, "total page number is", pageNumber, "grade =", grade)

}
