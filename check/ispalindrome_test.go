package palindrome

import(
	"testing"
)

func TestIspalindrome(t *testing.T){
	got := Ispalindrome("")
	want := true

	if got != want {
		t.Fatalf("Faulty palindrome checking algorithm")
	}
}
