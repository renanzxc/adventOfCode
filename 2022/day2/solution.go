package day2

import (
	"strings"
)

type GuideLineResut string

func (g GuideLineResut) Points() int64 {
	switch g {
	case Draw:
		return 3
	case Won:
		return 6
	}

	return 0
}

const (
	Won  GuideLineResut = "won"
	Lost GuideLineResut = "lost"
	Draw GuideLineResut = "draw"
)

type ChooseObject interface {
	ResultCompareObject(chooseObject ChooseObject) GuideLineResut
	GetChooseObjectTo(expectedGuideLineResut GuideLineResut) ChooseObject
	YourPoints() int64
}

type Rock struct {
}

func (r Rock) ResultCompareObject(chooseObject ChooseObject) GuideLineResut {
	switch chooseObject.(type) {
	case *Rock:
		return Draw
	case *Paper:
		return Lost
	}

	return Won
}

func (r Rock) YourPoints() int64 {
	return 1
}

func (r Rock) GetChooseObjectTo(expectedGuideLineResut GuideLineResut) ChooseObject {
	switch expectedGuideLineResut {
	case Won:
		return new(Paper)
	case Draw:
		return new(Rock)
	}

	return new(Scissors)
}

type Paper struct {
}

func (p Paper) ResultCompareObject(chooseObject ChooseObject) GuideLineResut {
	switch chooseObject.(type) {
	case *Rock:
		return Won
	case *Paper:
		return Draw
	}

	return Lost
}

func (p Paper) YourPoints() int64 {
	return 2
}

func (p Paper) GetChooseObjectTo(expectedGuideLineResut GuideLineResut) ChooseObject {
	switch expectedGuideLineResut {
	case Won:
		return new(Scissors)
	case Draw:
		return new(Paper)
	}

	return new(Rock)
}

type Scissors struct {
}

func (s Scissors) ResultCompareObject(chooseObject ChooseObject) GuideLineResut {
	switch chooseObject.(type) {
	case *Rock:
		return Lost
	case *Paper:
		return Won
	}

	return Draw
}

func (s Scissors) YourPoints() int64 {
	return 3
}

func (s Scissors) GetChooseObjectTo(expectedGuideLineResut GuideLineResut) ChooseObject {
	switch expectedGuideLineResut {
	case Won:
		return new(Rock)
	case Draw:
		return new(Scissors)
	}

	return new(Paper)
}

func newChooseObject(choose string) ChooseObject {
	switch choose {
	case "A", "X":
		return new(Rock)
	case "B", "Y":
		return new(Paper)
	}

	return new(Scissors)
}

func newGuideLineResult(resultIn string) GuideLineResut {
	switch resultIn {
	case "X":
		return Lost
	case "Y":
		return Draw
	}

	return Won
}

func Solution1(input []string) int64 {
	var totalPoints int64
	for _, strategyGuideLine := range input {
		var (
			strategyGuideLineList = strings.Split(strategyGuideLine, " ")
			opponentChoose        = newChooseObject(strategyGuideLineList[0])
			yourChoose            = newChooseObject(strategyGuideLineList[1])
			yourPoints            = yourChoose.YourPoints()
		)

		totalPoints += yourPoints + yourChoose.ResultCompareObject(opponentChoose).Points()
	}
	return totalPoints
}

func Solution2(input []string) int64 {
	var totalPoints int64
	for _, strategyGuideLine := range input {
		var (
			strategyGuideLineList = strings.Split(strategyGuideLine, " ")
			opponentChoose        = newChooseObject(strategyGuideLineList[0])
			needsToEnd            = newGuideLineResult(strategyGuideLineList[1])
		)

		yourChoose := opponentChoose.GetChooseObjectTo(needsToEnd)
		totalPoints += yourChoose.YourPoints() + needsToEnd.Points()
	}
	return totalPoints
}
