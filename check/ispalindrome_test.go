package palindrome

import (
	"testing"
)

func TestIspalindrome(t *testing.T) {
	got := Ispalindrome("0P")
	want := false

	if got != want {
		t.Fatalf("Fault palindrome checking algorithm")
	}
}
