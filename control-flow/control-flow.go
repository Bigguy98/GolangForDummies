package main

import (
	"fmt"
)

func waysOfLoop() {

	fmt.Println("***************** SHORT WAY ******************")
	i := 0
	for i < 10 {
		fmt.Println("couting", i)
		i++
	}

	fmt.Println("***************** NORMAL WAY ******************")
	for j := 0; j < 10; j++ {
		fmt.Println("couting", j)
	}

	fmt.Println("***************** IN RANGE WAY ******************")
	// _ will hold the index of iteration
	numbers := []int{1, 2, 3, 4}
	for _, number := range numbers {
		fmt.Println("couting", number)
	}
}

func main() {
	deferFunc := func() {
		fmt.Println("\"Defer\" define a function to execute when surrounding function complete.")
		fmt.Println("So this line will be printed at the end!")
	}
	defer deferFunc()

	waysOfLoop()

}
