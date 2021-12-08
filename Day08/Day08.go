package Day08

import (
	"fmt"
	"strings"
)

func solvePt1(inputLines []string) {

	count := 0
	for _, singleLine := range inputLines {
		firstSplit := strings.Split(singleLine, " | ")
		displayedOut := firstSplit[1]
		dispNumbers := strings.Split(displayedOut, " ")
		for _, singleDispNumb := range dispNumbers {
			switch len(singleDispNumb) {
			case 2:
				count++
			case 3:
				count++
			case 4:
				count++
			case 7:
				count++
			}
		}
	}
	fmt.Printf("Found: %v\n", count)

}

func solvePt2(inputLines []string) {

	sum := 0

	for _, singleLine := range inputLines {
		activeMap := make(map[int][]bool)
		firstSplit := strings.Split(singleLine, " | ")
		keyPart := firstSplit[0]
		keyNumbers := strings.Split(keyPart, " ")
		unknownKeyNumbs := make([][]bool, 0)
		for _, singleDispNumb := range keyNumbers {
			switch len(singleDispNumb) {
			case 2:
				activeMap[1] = getBoolArrayFromNumber(singleDispNumb)
			case 3:
				activeMap[7] = getBoolArrayFromNumber(singleDispNumb)
			case 4:
				activeMap[4] = getBoolArrayFromNumber(singleDispNumb)
			case 7:
				activeMap[8] = getBoolArrayFromNumber(singleDispNumb)
			default:
				unknownKeyNumbs = append(unknownKeyNumbs, getBoolArrayFromNumber(singleDispNumb))
			}
		}
		findFivePieceNumbers(&activeMap, &unknownKeyNumbs)
		findSixPieceNumbers(&activeMap, &unknownKeyNumbs)

		displayedNumbs := strings.Split(firstSplit[1], " ")
		newNum := 0
		for i := 0; i < len(displayedNumbs); i++ {
			thisNum := getNumFromMap(activeMap, displayedNumbs[i])
			newNum = newNum*10 + thisNum
			//fmt.Printf("%v\n", thisNum)
		}
		sum += newNum
		fmt.Printf("%v\n", newNum)
	}
	fmt.Printf("%v\n", sum)
}

func getNumFromMap(activeMap map[int][]bool, s string) int {
	targetBool := getBoolArrayFromNumber(s)

	//this might be overkill
	// check that the same number of segments are lit
	// check that anding does not change number of segments
	for k, v := range activeMap {
		if countTrue(v) == countTrue(targetBool) {
			if countTrue(andBoolArrays(v, targetBool)) == countTrue(v) {
				return k
			}
		}
	}
	//for k, v := range activeMap {
	//	if countTrue(andBoolArrays(v, targetBool)) == countTrue(v) {
	//		return k
	//	}
	//}
	return -1

}

func countTrue(boolArr []bool) int {
	count := 0
	for _, v := range boolArr {
		if v {
			count++
		}
	}
	return count
}

func findFivePieceNumbers(activeMap *map[int][]bool, unknownKeyNumbs *[][]bool) {
	//	find BD
	bdOnly := make([]bool, 7)
	//copy(bdOnly,(*activeMap)[4])
	//mask 4 with 7
	for i := 0; i < 7; i++ {
		if (*activeMap)[4][i] && !(*activeMap)[7][i] {
			bdOnly[i] = true
		}
	}

	//split out the 5 piece items, 2,3,5
	onlyFivePiece := make([][]bool, 0)
	newUnknownKeyNumbs := make([][]bool, 0)
	for _, val := range *unknownKeyNumbs {
		if countTrue(val) == 5 {
			onlyFivePiece = append(onlyFivePiece, val)
		} else {
			newUnknownKeyNumbs = append(newUnknownKeyNumbs, val)
		}
	}

	//find 5
	onlyTwoAndThree := make([][]bool, 0)
	for _, val := range onlyFivePiece {
		if countTrue(andBoolArrays(bdOnly, val)) == 2 {
			(*activeMap)[5] = val
		} else {
			onlyTwoAndThree = append(onlyTwoAndThree, val)
		}
	}
	for _, val := range onlyTwoAndThree {
		if countTrue(andBoolArrays(((*activeMap)[1]), val)) == 2 {
			(*activeMap)[3] = val
		} else {
			(*activeMap)[2] = val
		}
	}

	*unknownKeyNumbs = newUnknownKeyNumbs

}

func findSixPieceNumbers(activeMap *map[int][]bool, unknownKeyNumbers *[][]bool) {

	onlyNineAndSix := make([][]bool, 0)
	for _, val := range *unknownKeyNumbers {
		if countTrue(andBoolArrays((*activeMap)[4], val)) == 3 && countTrue(andBoolArrays((*activeMap)[1], val)) == 2 {
			(*activeMap)[0] = val
		} else {
			onlyNineAndSix = append(onlyNineAndSix, val)
		}
	}
	for _, val := range onlyNineAndSix {
		if countTrue(andBoolArrays((*activeMap)[7], val)) == 3 {
			(*activeMap)[9] = val
		} else {
			(*activeMap)[6] = val
		}
	}
}

func andBoolArrays(arr1, arr2 []bool) []bool {
	min := len(arr1)
	if len(arr2) < min {
		min = len(arr2)
	}
	returnArr := make([]bool, min)
	for i, _ := range returnArr {
		if arr1[i] && arr2[i] {
			returnArr[i] = true
		}
	}
	return returnArr
}

func getBoolArrayFromNumber(numbString string) []bool {
	returnArr := make([]bool, 7)
	numbArr := []rune(numbString)
	// a = 97 ...
	for _, val := range numbArr {
		returnArr[103-val] = true
	}

	return returnArr
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
