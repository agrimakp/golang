package main

import "fmt"

type Comic struct {
	title      string
	publisher  string
	writer     string
	artist     string
	year       int
	pageNumber int
	grade      float32
}

func main() {
	book1 := Comic{
		title:      "Mr. GoToSleep",
		publisher:  "DizzyBooks Publishing Inc.",
		writer:     "Tracey Hatchet",
		artist:     "Jewel",
		year:       1997,
		pageNumber: 0,
		grade:      0,
	}
	fmt.Println(book1.title, "written by", book1.writer, "drawn by", book1.artist, "Published by", book1.publisher, "in", book1.year, "total page number is", book1.pageNumber, "grade =", book1.grade)

	book2 := Comic{
		title:      "Mr. GoToSleep",
		publisher:  "DizzyBooks Publishing Inc.",
		writer:     "Tracey Hatchet",
		artist:     "Jewel",
		year:       1997,
		pageNumber: 0,
		grade:      0,
	}
	fmt.Println(book1.title, "written by", book1.writer, "drawn by", book1.artist, "Published by", book1.publisher, "in", book1.year, "total page number is", book1.pageNumber, "grade =", book1.grade)

	// title = "Epic Vol. 1"
	// publisher = "DizzyBooks Publishing Inc."
	// writer = "Ryan N. Shawn."
	// artist = "Phoebe Paperclips"
	// year = 2013
	// pageNumber = 160
	// grade = 9.0
	// fmt.Println(title, "written by", writer, "drawn by", artist, "Published by", publisher, "in", year, "total page number is", pageNumber, "grade =", grade)

}
