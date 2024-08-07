package word2

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	n := len(s)
	letters := make([]rune, n)
	for idx, r := range s {
		if unicode.IsLetter(r) {
			letters[idx] = unicode.ToLower(r)
		}
	}
	//fmt.Println(letters)
	for i := range letters {
		if letters[i] != letters[n-1-i] {
			return false
		}
	}
	return true
}
