package stringsXXX123

// name of package must same as containing folder
func ReverseString(str string) string {
	runes := []rune(str)
	length := len(runes)
	reversedRunes := make([]rune, length)
	for i := 0; i < length; i++ {
		reversedRunes[length-1-i] = runes[i]
	}
	return string(reversedRunes)
}
