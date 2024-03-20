package utils

import "strings"

// reverseString переворачивает строку
func ReverseString(s string) string {
	chars := strings.Split(s, "")
	reversedChars := make([]string, len(chars))
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		reversedChars[i], reversedChars[j] = chars[j], chars[i]
	}
	return strings.Join(reversedChars, "")
}
