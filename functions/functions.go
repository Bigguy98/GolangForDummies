package main

import (
	"fmt"
)

// varidic function
func calSum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// named return value
func sayHi() (x, y string) {
	x = "Hello world"
	y = "I'll shake the world"
	// before return, we assign value to variable for increasing readability
	return x, y
}

// recursive function
func feedMe(portion int, eaten int) {
	eaten += portion
	fmt.Println("I've just ate: ", eaten, ". I need more!")
	if eaten >= 5 {
		fmt.Println("Ok, I'm full.")
		return
	}
	feedMe(portion, eaten)
}

func showInfo(getName func() string) {
	fmt.Println("Hello", getName(), ". Have a good day!")
}

func main() {
	fmt.Println("***************** VARIDIC FUNCTION ******************")
	fmt.Println(calSum(100, 2999, 12, 11, 1))

	fmt.Println("***************** NAMED RETURN VALUE ******************")
	fmt.Println(sayHi)

	fmt.Println("***************** RECURSIVE FUNCTION ******************")
	feedMe(1, 0)

	fmt.Println("***************** PASSING FUNCTION AS PARAMETER ******************")
	getName := func() string {
		return "bugsmaker"
	}
	showInfo(getName)

}
