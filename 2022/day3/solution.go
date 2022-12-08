package day3

import "fmt"

const (
	numUpperCaseA = int('A')
	numLowerCaseA = int('a')
)

func Solution1(input []string) {
	var sum int
	for _, line := range input {
		equalsLetters := intersections(line[:len(line)/2], line[len(line)/2:])

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
		equalsLetters := intersections(input[ii], input[ii+1], input[ii+2])

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

func intersections(rucksacks ...string) (allIntersections []string) {
	var (
		intersectionsMap = make(map[string][]bool)
		rucksacksNum     = len(rucksacks)
	)

	for ii := range rucksacks {
		for _, letter := range rucksacks[ii] {

			if _, ok := intersectionsMap[string(letter)]; !ok {
				intersectionsMap[string(letter)] = make([]bool, rucksacksNum)
			}

			intersectionsMap[string(letter)][ii] = true
		}
	}

	for letter, rucksackHasLetter := range intersectionsMap {
		var allHas = true

		for _, hasLetter := range rucksackHasLetter {
			allHas = allHas && hasLetter
		}
		if allHas {
			allIntersections = append(allIntersections, letter)
		}
	}

	return
}

func gerNumLetter(letter string) int {
	numLetter := int([]rune(letter)[0])
	if numUpperCaseA <= numLetter && numLowerCaseA > numLetter {
		return numLetter - 38
	}
	return numLetter - 96
}
