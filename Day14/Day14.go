package Day14

import (
	"fmt"
	"sort"
	"strings"
)

func solveToSteps(inputLines []string, steps int) {

	polymerArray := inputLines[0]

	insertMap := make(map[string]string)
	currentPairs := make(map[string]int)
	currentCount := make(map[string]int)

	for _, val := range inputLines[2:] {
		insArray := strings.Split(val, " -> ")
		insertMap[insArray[0]] = insArray[1]
		currentCount[insArray[1]] = 0
	}

	for x := 0; x < len(polymerArray)-1; x++ {
		thisPair := polymerArray[x : x+2]
		if _, ok := currentPairs[thisPair]; ok {
			currentPairs[thisPair]++
		} else {
			currentPairs[thisPair] = 1
		}
	}

	for x := 0; x < len(polymerArray); x++ {
		thisSingle := polymerArray[x : x+1]
		if _, ok := currentCount[thisSingle]; ok {
			currentCount[thisSingle]++
		} else {
			currentCount[thisSingle] = 1
		}
	}

	for x := 0; x < steps; x++ {
		nextPairs := make(map[string]int)
		for key, _ := range insertMap {
			nextPairs[key] = 0
		}
		for key, val := range currentPairs {
			insertChar, _ := insertMap[key]
			currentCount[insertChar] += val
			//	add two new pair combos to nextPairs
			newPair := key[:1] + insertChar
			nextPairs[newPair]+=val
			newPair = insertChar + key[1:]
			nextPairs[newPair]+=val
		}
		currentPairs=nextPairs
		fmt.Printf("%v\n",currentCount)
	}
	counts:=make([]int,0)
	for _,val:=range currentCount{
		counts= append(counts, val)
	}
	sort.Ints(counts)

	fmt.Printf("%v\n",counts[len(counts)-1]-counts[0])

}

func solvePt1(inputLines []string) {
	solveToSteps(inputLines, 10)
}

func solvePt2(inputLines []string) {
	solveToSteps(inputLines, 40)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
