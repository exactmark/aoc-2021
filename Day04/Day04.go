package Day04

import (
	"fmt"
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

func createBoard() *bingoBoard {
	thisBoard := bingoBoard{
		heldNumbers:    make(map[int]coord),
		occupiedSpaces: make(map[coord]empty),
	}
	return &thisBoard
}

func (b *bingoBoard) populateBoard(inputLines []string) {
	for y := 0; y < len(inputLines); y++ {
		xPtr := 0

		for x := 0; xPtr < 5; x += 2 {
			thisCoord := coord{
				x: xPtr,
				y: y,
			}
			newNum, _ := strconv.Atoi(
				strings.TrimSpace(
					inputLines[y][x+xPtr : x+xPtr+2]))
			b.heldNumbers[newNum] = thisCoord
			xPtr++
		}
	}
}

func (b *bingoBoard) isWinner() bool {
	//Crackers, diagonals don't count.
	////check for diagonal wins
	//hasDiag := true
	//for i := 0; i < 5; i++ {
	//	testCoord := coord{
	//		x: i,
	//		y: i,
	//	}
	//	if _, ok := b.occupiedSpaces[testCoord]; !ok {
	//		hasDiag = false
	//		break
	//	}
	//}
	//if hasDiag {
	//	return true
	//}
	//hasDiag=true
	//for i := 0; i < 5; i++ {
	//	testCoord := coord{
	//		x: i,
	//		y: 4 - i,
	//	}
	//	if _, ok := b.occupiedSpaces[testCoord]; !ok {
	//		hasDiag = false
	//		break
	//	}
	//}
	//if hasDiag {
	//	return true
	//}

	//check for non-diagonal wins
	yCounts := make([]int, 5)
	xCounts := make([]int, 5)
	for k, _ := range b.occupiedSpaces {
		xCounts[k.x]++
		yCounts[k.y]++
		if xCounts[k.x] == 5 || yCounts[k.y] == 5 {
			return true
		}
	}

	return false
}

func (b *bingoBoard)processNumber(inputNum int){
	if val,ok:= b.heldNumbers[inputNum];ok{
		b.occupiedSpaces[val]=empty{}
		delete(b.heldNumbers,inputNum)
	}
}

func (b *bingoBoard)getRemainingSum()int{
	sum:=0
	for k,_:=range b.heldNumbers{
		sum+=k
	}
	return sum
}

func (b *bingoBoard) toString() string {
	returnString := ""
	returnString += fmt.Sprintf("remaining: %v\n", len(b.heldNumbers))
	returnString += fmt.Sprintf("held: %v\n", b.heldNumbers)
	returnString += fmt.Sprintf("occupied: %v\n", b.occupiedSpaces)
	return returnString
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
		nextBoard := createBoard()
		nextBoard.populateBoard(inputLines[i+1 : i+6])
		allBoards = append(allBoards, nextBoard)
	}

	for i:=0;i<len(randomNumbers);i++{
		nextNumber:=randomNumbers[i]
		for _,board:=range allBoards{
			board.processNumber(nextNumber)
			if board.isWinner(){
				fmt.Printf(board.toString())
				i=len(randomNumbers)
				fmt.Printf("remSum:%v\n",board.getRemainingSum())
				fmt.Printf("winningScore:%v\n",board.getRemainingSum()*nextNumber)
				break
			}
		}
	}


	//fmt.Printf("%v\n", randomNumbers)
	//fmt.Printf("%v\n", allBoards)
	//for _, v := range allBoards {
	//	fmt.Printf(v.toString())
	//	fmt.Printf("%v\n", v.isWinner())
	//}
}

func solvePt2(inputLines []string) {

	var randomNumbers []int

	lineSplit := strings.Split(inputLines[0], ",")
	for _, v := range lineSplit {
		singleNum, _ := strconv.Atoi(v)
		randomNumbers = append(randomNumbers, singleNum)
	}

	var allBoards []*bingoBoard

	for i := 1; i < len(inputLines); i += 6 {
		nextBoard := createBoard()
		nextBoard.populateBoard(inputLines[i+1 : i+6])
		allBoards = append(allBoards, nextBoard)
	}

	remainingBoards:=len(allBoards)

	for i:=0;i<len(randomNumbers);i++{
		nextNumber:=randomNumbers[i]
		for boardPtr,board:=range allBoards{
			if board!=nil {
				board.processNumber(nextNumber)
				if board.isWinner() {
					remainingBoards-=1
					if remainingBoards==0{
						fmt.Printf(board.toString())
						i=len(randomNumbers)
						fmt.Printf("remSum:%v\n",board.getRemainingSum())
						fmt.Printf("winningScore:%v\n",board.getRemainingSum()*nextNumber)
						break
					}
					allBoards[boardPtr] = nil
				}
			}
		}
	}
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
