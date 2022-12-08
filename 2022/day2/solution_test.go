package day2

import (
	"adventofcode/utils"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay2(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	input := utils.ReadFile(filepath.Join(path, "input.txt"))
	t.Run("Solution1", func(t *testing.T) {
		var expected int64 = 10941
		if result := Solution1(input); expected != result {
			t.Fatalf(`Expected:%d Result:%d`, expected, result)
		}
	})

	t.Run("Solution2", func(t *testing.T) {
		var expected int64 = 13071
		if result := Solution2(input); expected != result {
			t.Fatalf(`Expected:%d Result:%d`, expected, result)
		}
	})
}
