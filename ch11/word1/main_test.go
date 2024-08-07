package word1

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("ooglelgoo") {
		t.Error(`IsPalindrom("oogleloo") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = true`)
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("retarded") {
		t.Error(`IsPalindrome("retarded")= false`)
	}
}
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
