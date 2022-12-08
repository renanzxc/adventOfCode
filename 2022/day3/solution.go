package day3

import (
	"adventofcode/utils"
	"fmt"
)

const (
	numUpperCaseA = int('A')
	numLowerCaseA = int('a')
)

func Solution1(input []string) {
	var sum int
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
	fmt.Println(sum)
}

func Solution2(input []string) {
	var sum int
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
	fmt.Println(sum)
}

func gerNumLetter(letter rune) int {
	numLetter := int(letter)
	if numUpperCaseA <= numLetter && numLowerCaseA > numLetter {
		return numLetter - 38
	}
	return numLetter - 96
}
