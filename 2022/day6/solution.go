package day6

func Solution1(input []string) any {
	return getNumCharactersBeforeMark(input[0], 4)
}

func Solution2(input []string) any {
	return getNumCharactersBeforeMark(input[0], 14)
}

func getNumCharactersBeforeMark(line string, numUniqueCharactersToMark int) int {
	var (
		lenLine   = len(line)
		startFrom int
	)

	for {
		var (
			mapLetters                  = make(map[string]bool)
			hasRepeatedLetterOnInterval bool
		)

		for _, letter := range line[startFrom : startFrom+numUniqueCharactersToMark] {
			ok := mapLetters[string(letter)]
			if ok {
				hasRepeatedLetterOnInterval = true
				break
			}
			mapLetters[string(letter)] = true
		}

		if !hasRepeatedLetterOnInterval {
			break
		}

		startFrom++
		if startFrom+numUniqueCharactersToMark > lenLine {
			break
		}
	}

	return startFrom + numUniqueCharactersToMark
}
