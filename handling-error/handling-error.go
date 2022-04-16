package main

import (
	"errors"
	"fmt"
)

func devideHalf(x int) (int, error) {
	if x%2 == 0 {
		return x / 2, nil
	}
	return -1, errors.New("This number can not divide to 2")
}

func main() {
	fmt.Println("***************** HANDLING ERROR EXAMPLE ************************")
	err := errors.New("Something went wrong!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("***************** CREATEING ERROR FORMAT WITH FMT ************************")
	err2 := fmt.Errorf("The player %v with role %v has been quit!", "bugsmaker", "admin")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("***************** FUNCTION RETURN ERROR EXAMPLE ************************")
	x2, err := devideHalf(23)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Devide value: ", x2)
	}

	fmt.Println("***************** PANIC LIB EXAMPLE ************************")
	fmt.Println("Something went wrong!")
	panic("OMG, what should we do?")
}
