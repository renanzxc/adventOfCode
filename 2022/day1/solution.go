package day1

import (
	"adventofcode/utils"
	"log"
	"strconv"
)

func Solution1(input []string) any {
	return getMostCaloriesElfs(input, 1)[0]
}

func Solution2(input []string) any {
	return utils.Sum(getMostCaloriesElfs(input, 3)...)
}

func getMostCaloriesElfs(input []string, numElfs int64) []int64 {
	var (
		maxCaloriesElfs   = make([]int64, numElfs)
		actualCaloriesElf int64
	)
	for _, caloriesStr := range input {
		if caloriesStr == "" {
			for elfII := range maxCaloriesElfs {
				if actualCaloriesElf > maxCaloriesElfs[elfII] {
					maxCaloriesElfs = utils.InsertArray(maxCaloriesElfs, elfII, actualCaloriesElf)[:numElfs]
					break
				}
			}
			actualCaloriesElf = 0

			continue
		}

		calories, err := strconv.ParseInt(caloriesStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		actualCaloriesElf += calories
	}

	return maxCaloriesElfs
}
