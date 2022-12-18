package day6

import (
	"adventofcode/utils"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay6(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	input := utils.ReadFile(filepath.Join(path, "input.txt"))

	t.Run("Solution1", func(t *testing.T) {
		t.Run("Example", func(t *testing.T) {
			var expected any = 5
			if result := Solution1([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}); expected != result {
				t.Fatalf(`Expected:%s Result:%s`, expected, result)
			}
		})
		t.Run("Input", func(t *testing.T) {
			var expected any = 1623
			if result := Solution1(input); expected != result {
				t.Fatalf(`Expected:%s Result:%s`, expected, result)
			}
		})
	})

	t.Run("Solution2", func(t *testing.T) {
		t.Run("Example", func(t *testing.T) {
			var expected any = 23
			if result := Solution2([]string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}); expected != result {
				t.Fatalf(`Expected:%s Result:%s`, expected, result)
			}
		})
		//t.Run("Input", func(t *testing.T) {
		//	var expected any = 1623
		//	if result := Solution1(input); expected != result {
		//		t.Fatalf(`Expected:%s Result:%s`, expected, result)
		//	}
		//})
	})
}
