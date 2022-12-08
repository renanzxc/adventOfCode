package day4

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solution1(input []string) {
	var count int
	for _, line := range input {
		initElfSections1, endElfSections1, initElfSections2, endElfSections2 := extractElfSections(line)

		if (initElfSections1 <= initElfSections2 && endElfSections1 >= endElfSections2) ||
			(initElfSections2 <= initElfSections1 && endElfSections2 >= endElfSections1) {
			count++
		}
	}
	fmt.Println(count)
}

func Solution2(input []string) {
	var count int
	for _, line := range input {
		initElfSections1, endElfSections1, initElfSections2, endElfSections2 := extractElfSections(line)

		var elfSections1Range, elfSections2Range []int64
		for ii := initElfSections1; ii <= endElfSections1; ii++ {
			elfSections1Range = append(elfSections1Range, ii)
		}
		for ii := initElfSections2; ii <= endElfSections2; ii++ {
			elfSections2Range = append(elfSections2Range, ii)
		}

		if len(utils.Intersections(elfSections1Range, elfSections2Range)) > 0 {
			count++
		}
	}
	fmt.Println(count)
}

func extractElfSections(line string) (initElfSections1, endElfSections1, initElfSections2, endElfSections2 int64) {
	var (
		elfSections                = strings.Split(line, ",")
		elfSections1, elfSections2 = strings.Split(elfSections[0], "-"), strings.Split(elfSections[1], "-")
		err                        error
	)

	initElfSections1, err = strconv.ParseInt(elfSections1[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	endElfSections1, err = strconv.ParseInt(elfSections1[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	initElfSections2, err = strconv.ParseInt(elfSections2[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	endElfSections2, err = strconv.ParseInt(elfSections2[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return
}
