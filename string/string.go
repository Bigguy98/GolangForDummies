package main

import (
	"bytes"
	"fmt"
	"strings"
)

// using buffer to concat string
func concatStringWithBuff() {
	var buffer bytes.Buffer
	for i := 0; i <= 2; i++ {
		buffer.WriteString("tungdz\n")
	}
	fmt.Println(buffer.String())
}

func main() {
	s1 := `After a backslash, certain single
	character escapes represent
	special values\n
	n is a line feed or new line\t
	t is a tab`
	fmt.Println(s1)
	// with the "" you can not declare a multiline string
	// you have to use + to connect lines
	s2 := "After After a backslash, certain single character escapes represent special " +
		"values\nn is a line feed or new line\tt is a tab"
	fmt.Println(s2)

	fmt.Println("**************** CONCAT STRING WITH BUFFER *************")
	concatStringWithBuff()

	fmt.Println("**************** WORKING WITH STRING LIBS *************")
	fmt.Println(strings.Replace(s2, "After", "Before", 3))
}
