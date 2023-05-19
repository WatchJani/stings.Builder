package main

import (
	"fmt"
	"strings"
)

func ConCat(data []string) {
	var res string

	for i := 0; i < len(data); i++ {
		res += data[i]
	}
}

func BetterString(data []string) {
	var res strings.Builder

	res.Grow(15) //capacity

	for i := 0; i < 4; i++ {
		res.WriteString(data[i])
	}
}

func CombineStrings(list []string) string {
	var combineStrings strings.Builder

	for i := 0; i < len(list); i++ {
		combineStrings.WriteString(list[i])

		if i < len(list)-1 {
			combineStrings.WriteByte(' ')
		}
	}

	return combineStrings.String()
}

func RemoveDuplicates(input string) string {
	var builder strings.Builder
	seen := make(map[rune]bool)

	for _, char := range input {
		if !seen[char] {
			builder.WriteRune(char)
			seen[char] = true
		}
	}

	return builder.String()
}

func main() {
	myData := []string{"my", "super", "data", "test"}

	// ConCat(myData)
	// BetterString(myData)

	fmt.Println(CombineStrings(myData))
}
