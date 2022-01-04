package main

import "fmt"

var A = '0'
var END = '9'

func ValidISBN10(isbns string) bool {
	sum := 0
	isbn := []rune(isbns)
	for i := 0; i < len(isbn); i++ {
		if isbn[i] < A || isbn[i] > END {
			if isbn[i] != 'X' {
				return false
			}
			sum += 10 * (i + 1)
			continue
		}
		r := int(isbn[i] - A)
		sum += r * (i + 1)

	}

	return sum%11 == 0
}

func main() {
	fmt.Println(ValidISBN10("048665088X"))
}
