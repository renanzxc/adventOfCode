package day1

import (
	"adventofcode/utils"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay1(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	input := utils.ReadFile(filepath.Join(path, "input.txt"))
	t.Run("Solution1", func(t *testing.T) {
		var expected int64 = 69206
		if result := Solution1(input); expected != result {
			t.Fatalf(`Expected:%d Result:%d`, expected, result)
		}
	})

	t.Run("Solution2", func(t *testing.T) {
		var expected int64 = 197400
		if result := Solution2(input); expected != result {
			t.Fatalf(`Expected:%d Result:%d`, expected, result)
		}
	})
}
