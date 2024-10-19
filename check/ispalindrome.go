package palindrome

import (
	"unicode"
)

func Ispalindrome(source string) bool {
	leftPointer := 0
	lenSource := len(source)

	rightPointer := lenSource - 1

	result := true
	var leftChar rune
	var rightChar rune
	for {
		if leftPointer >= rightPointer {
			break
		}
		for {
			leftChar = unicode.ToLower(rune(source[leftPointer]))
			leftPointer++
			if unicode.IsLetter(leftChar) || unicode.IsDigit(leftChar) {
				break
			}
			if leftPointer >= lenSource {
				leftChar = rune(0)
				break
			}
		}
		for {
			rightChar = unicode.ToLower(rune(source[rightPointer]))
			rightPointer--
			if unicode.IsLetter(rightChar) || unicode.IsDigit(rightChar) {
				break
			}
			if rightPointer < 0 {
				rightChar = rune(0)
				break
			}
		}
		if leftChar != rightChar {
			result = false
			break
		}

	}

	return result
}
