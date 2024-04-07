package reverse_string

import "fmt"

func ReverseString(inputString string) string {
	runes := []rune(inputString)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
	return string(runes)
}
