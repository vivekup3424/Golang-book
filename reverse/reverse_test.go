package reverse

import (
	"testing"

	"github.com/kernignan-book/reverse"
)

type Word struct {
	in, want string
}

func TestString(t *testing.T) {
	v := []Word{
		{"Hello", "olleH"},
		{"App le", "el ppA"},
	}
	for _, c := range v {
		got := (c.in)
		result := reverse.String(got)
		if result != c.want {
			t.Errorf("Assertion Error: res(%s) != want(%s)",
				&result, c.want)
		}

	}
}
