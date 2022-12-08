package utils

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

func ReadFile(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return
}

func InsertArray[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func Sum[T Numeric](num ...T) (total T) {
	for ii := range num {
		total += num[ii]
	}

	return
}

func Intersections[T comparable](itens ...[]T) (allIntersections []T) {
	var (
		intersectionsMap = make(map[T][]bool)
		itensNum         = len(itens)
	)

	for ii := range itens {
		for _, item := range itens[ii] {
			if _, ok := intersectionsMap[item]; !ok {
				intersectionsMap[item] = make([]bool, itensNum)
			}

			intersectionsMap[item][ii] = true
		}
	}

	for item, hasItens := range intersectionsMap {
		var allHas = true

		for _, hasItem := range hasItens {
			allHas = allHas && hasItem
		}
		if allHas {
			allIntersections = append(allIntersections, item)
		}
	}

	return
}
