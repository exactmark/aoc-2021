package Day01

import (
	"fmt"

	"strconv"
)

func countAscents(inputInts []int) int {

	count:=0

	for x:=1; x<len(inputInts); x++ {
		if inputInts[x]>inputInts[x-1]{
			count++
		}
	}
	return count
}

func countTripleAscents(inputInts []int) int {

	count:=0

	for x:=3; x<len(inputInts); x++ {
		// x-1 and x-2 are common, so really just compare the following
		if inputInts[x]>inputInts[x-3]{
			count++
		}
	}
	return count
}

func solvePt1(inputLines []string) {
	var numList = make([]int, 0)
	for _, singleLine := range inputLines {
		thisNumber, _ := strconv.Atoi(singleLine)
		numList = append(numList, thisNumber)
	}
	numAscents := countAscents(numList)
	fmt.Printf("%v\n", numAscents)

}

func solvePt2(inputLines []string) {
	var numList = make([]int, 0)
	for _, singleLine := range inputLines {
		thisNumber, _ := strconv.Atoi(singleLine)
		numList = append(numList, thisNumber)
	}
	numAscents := countTripleAscents(numList)
	fmt.Printf("3s: %v\n", numAscents)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
