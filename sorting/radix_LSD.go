package sorting

import "fmt"

// SortRadixLSD sorts the strings in the input.
// Assumes that all strings have the same length
// Assumes that it only contains lowercase letters of the English alphabet
func SortRadixLSD(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	result := input
	for i := len(input[0]) - 1; i >= 0; i-- {
		result = sortUsingPosition(result, i)
	}

	return result
}

func sortUsingPosition(input []string, pos int) []string {
	if len(input) == 0 {
		return []string{}
	}
	if pos < 0 || pos >= len(input[0]) {
		panic(fmt.Sprintf("invalid position: %v", pos))
	}

	countPerRune := make([]int, 26)
	for _, s := range input {
		countPerRune[s[pos]-'a']++
	}

	startForRune := make([]int, 26)
	start := 0
	for i, c := range countPerRune {
		startForRune[i] = start
		start += c
	}

	result := make([]string, len(input))
	for _, s := range input {
		char := s[pos] - 'a'
		result[startForRune[char]] = s
		startForRune[char]++
	}

	return result
}
