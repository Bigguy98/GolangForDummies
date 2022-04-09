package main

import (
	"fmt"
)

/**
	it's not possible to add addition element to the array
**/
func showArrayExample() {
	var customers [3]string
	customers[0] = "bugsmaker"
	customers[1] = "tungdz"
	fmt.Println(customers[2])
	fmt.Println(customers)
}

/**
	slice can be consider as a wrapper of array, making it easier to work with
**/
func showSliceExample() {
	customers := make([]string, 2)
	customers[0] = "bugsmaker"
	customers[1] = "tungdz"
	fmt.Println(customers)
	customers = append(customers, "donal trump")
	fmt.Println(customers)
	customers = append(customers, "Elon Musk", "Da vinci", "Geogre Bush")
	fmt.Println(customers)
	// remove element at index 2
	customers = append(customers[:2], customers[3:]...)
	fmt.Println(customers)
	// copy slice
	newCustomer := make([]string, 2)
	copy(newCustomer, customers)
	fmt.Println(newCustomer)
}

func showMapExample() {
	playersScore := make(map[string]int)
	playersScore["Ronaldo"] = 100
	playersScore["Messi"] = 99
	playersScore["Benzema"] = 90
	fmt.Println(playersScore)
	delete(playersScore, "Messi")
	fmt.Println(playersScore)
}

func main() {
	fmt.Println("***************** ARRAY *******************")
	showArrayExample()
	fmt.Println("***************** SLICE *******************")
	showSliceExample()
	fmt.Println("***************** MAP *******************")
	showMapExample()
}
