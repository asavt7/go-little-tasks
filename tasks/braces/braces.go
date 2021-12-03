package braces

import "fmt"

//GenerateCorrectSequences generates slice of correct braces sequences of len 2n
func GenerateCorrectSequences(n int) []string {
	result := make([]string, 0)
	gen("", n, 0, 0, &result)
	return result
}

func gen(current string, maxBraces int, left int, right int, results *[]string) {
	if len(current) >= 2*maxBraces {
		fmt.Println(current)
		*results = append(*results, current)
		return
	}

	if left < maxBraces {
		gen(current+"(", maxBraces, left+1, right, results)
	}
	if right < left {
		gen(current+")", maxBraces, left, right+1, results)
	}

}
