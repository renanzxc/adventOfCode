package utils

import (
	"bufio"
	"log"
	"os"
)

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

type Numeric interface {
	~int64 | ~float64 | ~int | ~float32
}

func Sum[T Numeric](num ...T) (total T) {
	for ii := range num {
		total += num[ii]
	}

	return
}
