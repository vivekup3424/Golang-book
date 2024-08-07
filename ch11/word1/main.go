package word1

func IsPalindrome(s string) bool {
	n := len(s)
	for i := range s {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
