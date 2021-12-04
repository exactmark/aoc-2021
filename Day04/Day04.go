package Day04

import (
	"strconv"
	"strings"
)

type empty struct {
}

type coord struct {
	x int
	y int
}

type bingoBoard struct {
	heldNumbers    map[int]coord
	occupiedSpaces map[coord]empty
}

func generateBoard(inputLines []string)*bingoBoard{
	newBoard:=bingoBoard{
		heldNumbers:    nil,
		occupiedSpaces: nil,
	}

	return &newBoard
}

func solvePt1(inputLines []string) {

	var randomNumbers []int

	lineSplit := strings.Split(inputLines[0], ",")
	for _, v := range lineSplit {
		singleNum, _ := strconv.Atoi(v)
		randomNumbers = append(randomNumbers, singleNum)
	}

	var allBoards []*bingoBoard

	for i := 1; i < len(inputLines); i += 6 {
		nextBoard := generateBoard(inputLines[i+1 : i+6])
		allBoards= append(allBoards,nextBoard)
	}

}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
