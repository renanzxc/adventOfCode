package day3

import (
	"adventofcode/utils"
)

const (
	numUpperCaseA = int64('A')
	numLowerCaseA = int64('a')
)

func Solution1(input []string) int64 {
	var sum int64
	for _, line := range input {
		equalsLetters := utils.Intersections([]rune(line[:len(line)/2]), []rune(line[len(line)/2:]))

		var priority = gerNumLetter(equalsLetters[0])
		for _, letter := range equalsLetters[1:] {
			numLetter := gerNumLetter(letter)
			if priority > numLetter {
				priority = numLetter
			}
		}

		sum += priority

	}
	return sum
}

func Solution2(input []string) int64 {
	var sum int64
	for ii := 0; ii < len(input); ii++ {
		equalsLetters := utils.Intersections([]rune(input[ii]), []rune(input[ii+1]), []rune(input[ii+2]))

		var priority = gerNumLetter(equalsLetters[0])
		for _, letter := range equalsLetters[1:] {
			numLetter := gerNumLetter(letter)
			if priority > numLetter {
				priority = numLetter
			}
		}

		sum += priority
		ii += 2
	}
	return sum
}

func gerNumLetter(letter rune) int64 {
	numLetter := int64(letter)
	if numUpperCaseA <= numLetter && numLowerCaseA > numLetter {
		return numLetter - 38
	}
	return numLetter - 96
}
