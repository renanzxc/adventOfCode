package day5

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Solution1(input []string) any {
	return readStacksAndMovements(input)
}

func Solution2(input []string) any {
	return readStacksAndMovements2(input)
}

type Stack struct {
	Crates []Crate
}

type Crate string

func readStacksAndMovements(input []string) string{
	var (
		initLineMovements int
		stacks []Stack
	)
	for ii := range input{
		if input[ii] == ""{
			initLineMovements = ii +1
			break
		}
		var (
			actualStack = 0
			posi = 0
		)
		for _, char := range strings.Split(input[ii], ""){
			if posi == 3 {
				actualStack++
				posi = 0
				continue
			}

			var charStr = string(char)
			if isLetter(charStr) {
				if lastStack := len(stacks) - 1; lastStack < actualStack {
					stacks = append(stacks, make([]Stack, actualStack-lastStack)...)
				}

				stacks[actualStack].Crates = append([]Crate{Crate(charStr)},stacks[actualStack].Crates...)
			}

			posi++
		}
	}

	for ii := initLineMovements; ii < len(input); ii++ {
		lineSep := strings.Split(input[ii]," ")

		move, err :=strconv.ParseInt(lineSep[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		from, err :=strconv.ParseInt(lineSep[3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		to, err :=strconv.ParseInt(lineSep[5], 10, 64)
		if err != nil {
			log.Fatal(err)
		}


		// TODO: Alterar
		a := len(stacks[from-1].Crates)-int(move)-1
		for jj := len(stacks[from-1].Crates)-1; jj > a; jj--{
			stacks[to-1].Crates = append(stacks[to-1].Crates, stacks[from-1].Crates[jj])
			stacks[from-1].Crates = stacks[from-1].Crates[:jj]
		}
	}
	var letters string
	for ii := range stacks{
		letters += string(stacks[ii].Crates[len(stacks[ii].Crates)-1])
	}
	return letters
}

func readStacksAndMovements2(input []string) string{
	var (
		initLineMovements int
		stacks []Stack
	)
	for ii := range input{
		if input[ii] == ""{
			initLineMovements = ii +1
			break
		}
		var (
			actualStack = 0
			posi = 0
		)
		for _, char := range strings.Split(input[ii], ""){
			if posi == 3 {
				actualStack++
				posi = 0
				continue
			}

			var charStr = string(char)
			if isLetter(charStr) {
				if lastStack := len(stacks) - 1; lastStack < actualStack {
					stacks = append(stacks, make([]Stack, actualStack-lastStack)...)
				}

				stacks[actualStack].Crates = append([]Crate{Crate(charStr)},stacks[actualStack].Crates...)
			}

			posi++
		}
	}

	for ii := initLineMovements; ii < len(input); ii++ {
		lineSep := strings.Split(input[ii]," ")

		move, err :=strconv.ParseInt(lineSep[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		from, err :=strconv.ParseInt(lineSep[3], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		to, err :=strconv.ParseInt(lineSep[5], 10, 64)
		if err != nil {
			log.Fatal(err)
		}


		// TODO: Alterar
		stacks[to-1].Crates = append(stacks[to-1].Crates, stacks[from-1].Crates[len(stacks[from-1].Crates)-int(move):]...)
		stacks[from-1].Crates = stacks[from-1].Crates[:len(stacks[from-1].Crates)-int(move)]

	}
	var letters string
	for ii := range stacks{
		letters += string(stacks[ii].Crates[len(stacks[ii].Crates)-1])
	}
	return letters
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}