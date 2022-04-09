package main

import (
	"fmt"
)

// Movie is ...
type Movie struct {
	Name   string
	Rating float32
}

type Drink struct {
	Name string
	Ice  bool
}

func newMovie() Movie {
	return Movie{
		Name:   "God father",
		Rating: 7.88,
	}
}

func showStructExample() {
	var m1 Movie
	m1.Name = "Men in Black"
	m1.Rating = 6.22
	fmt.Println(m1)
	fmt.Printf("%+v\n", m1)

	m2 := new(Movie)
	m2.Name = "Private of Carribean"
	m2.Rating = 9.0
	fmt.Printf("%+v\n", m2)

	m3 := Movie{Name: "Citizen Kane", Rating: 10}
	fmt.Printf("%+v\n", m3)

	// omit the field names and assign values in the order they are declared
	m4 := Movie{"Citizen Kane", 10}
	fmt.Printf("%+v\n", m4)

	m5 := Movie{
		Name:   "OnePiece",
		Rating: 9.7,
	}
	fmt.Printf("%+v\n", m5)

	fmt.Printf("%+v\n", newMovie())

	if m4 == m5 {
		fmt.Println("Two movie a the same!")
	}

}

func assignValueVsAssignPointer() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	// assign value, creating copy of a and assign to b
	b := a
	a.Name = "Caffe"
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

	// assign pointer, a and c now point to the same memory address
	c := &a
	a.Name = "Capuchino"
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", *c)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", c)
}

func main() {
	fmt.Println("***************  SHOW STRUCT EXAMPLE ****************")
	showStructExample()
	fmt.Println("***************  ASSIGN POINTER VS ASSIGN VALUE ****************")
	assignValueVsAssignPointer()
}
