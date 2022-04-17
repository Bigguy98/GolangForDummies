package main

import (
	number "bugsmaker/mods/demo/numbers" // declare alias for package
	stringsXXX "bugsmaker/mods/demo/strings"
	greeting "bugsmaker/mods/demo/strings/gretting"
	"fmt"
)

/**
	lines in the import is just path to the package
	the package name must be the first line of package code
	or we can use alias
**/
func main() {
	fmt.Println("144 is prime number:", number.IsPrime(144))
	fmt.Println("reverse string of bugsmaker is:", stringsXXX.ReverseString("bugsmaker"))
	fmt.Println("greeting:", greeting.EveningText)
}
