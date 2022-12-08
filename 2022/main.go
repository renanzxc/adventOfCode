package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"adventofcode/day4"
	"adventofcode/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var days = map[string]func(input []string){
		"1-1": day1.Solution1,
		"1-2": day1.Solution2,
		"2-1": day2.Solution1,
		"2-2": day2.Solution2,
		"3-1": day3.Solution1,
		"3-2": day3.Solution2,
		"4-1": day4.Solution1,
		"4-2": day4.Solution2,
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var (
		dayIn       = os.Args[1]
		day         = strings.Split(dayIn, "-")[0]
		filepathDay = filepath.Join(path, fmt.Sprintf("day%s", day), "input.txt")
	)

	days[dayIn](utils.ReadFile(filepathDay))
}
