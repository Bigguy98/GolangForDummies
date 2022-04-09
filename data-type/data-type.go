package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func addition(a int, b int) int {
	return a - b
}

func sayHello(s string) string {
	return "Hello " + s
}

func showMemAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	fmt.Println(sayHello("bugsmaker!"))
	fmt.Println(100, 200)
	fmt.Println(addition(100, 200))

	var b bool
	fmt.Println(b)
	b = true
	fmt.Println(b)

	// int can be signed 32-bit integer or signed 64-bit integer.
	var i int = 100
	fmt.Println(i)

	var f32 float32 = 0.1112
	fmt.Println(f32)
	var f64 float64 = 0.19281983
	fmt.Println(f64)

	var s string = "Hello world@--xase123"
	fmt.Println(s)

	// When initializing an array, both the length and type of an array must be specified.
	var fruits [4]string
	fruits[0] = "Banana"
	fruits[1] = "Orange"
	fruits[2] = "Lemon"
	fruits[3] = "Watermelon"
	fmt.Println(fruits)

	// checking type of variables
	fmt.Println("================CHECKING TYPE OF VARIABLES================")
	fmt.Println(reflect.TypeOf(fruits))
	fmt.Println(reflect.TypeOf(f32))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))

	// convert between types
	fmt.Println("================CONVERT BETWEEN TYPES================")
	var boolStr string = strconv.FormatBool(b)
	fmt.Println(boolStr)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(boolStr))

	// shorthand variable decleration
	fmt.Println("================SHORTHAND VAR DECLERATION================")
	var s1, s2, s3 string = "foo", "bar", "foobar"
	fmt.Println(s1 + " " + s2 + " " + s3)
	// shorthand variable decleration for difference types
	var (
		dt1 string = "bugsmaker"
		dt2 int    = 2022
	)
	fmt.Println(dt1)
	fmt.Println(dt2)

	// initial value
	fmt.Println("================INITIAL VALUES ================")
	var initI int
	var initS string
	var initF float32
	var initB bool
	fmt.Printf("%v %q %v %v\n", initI, initS, initF, initB)
	if initS == "" {
		fmt.Println("initS has not been assigned a value and is zero valued")
	}

	// initial value
	fmt.Println("================SHORT VARIABLE FORM ================")
	sShort := "Hello World"
	fmt.Println(sShort)
	// pointer allocated a position in the computer’s memory, the sequence of number is in base 16 format
	fmt.Println("================POINTER ================")
	fmt.Println(&sShort)
	var pointerEx = 100
	fmt.Println(&pointerEx)
	showMemAddress(&pointerEx)
}
