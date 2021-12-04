package Day03

import (
	"fmt"
	"strconv"
)

func byteArrayToInt(bitArray []rune) int64 {
	intVal, _ := strconv.ParseInt(string(bitArray), 2, 64)
	return intVal
}

func solvePt1(inputLines []string) {
	countingArray := make([]int, len(inputLines[0]))
	for _, singleLine := range inputLines {
		charArray := []rune(singleLine)
		for k, v := range charArray {
			if v == '1' {
				countingArray[k] += 1
			}
		}

	}

	gammaBits := make([]rune, len(countingArray))
	epsilonBits := make([]rune, len(countingArray))
	for k, v := range countingArray {
		if float32(v)/float32(len(inputLines)) > 0.5 {
			gammaBits[k] = '1'
			epsilonBits[k] = '0'
		} else {
			gammaBits[k] = '0'
			epsilonBits[k] = '1'
		}
	}

	gammaInt := byteArrayToInt(gammaBits)
	epsilonInt := byteArrayToInt(epsilonBits)
	fmt.Printf("%v,%v\n", gammaBits, epsilonBits)
	fmt.Printf("%v,%v\n", gammaInt, epsilonInt)
	fmt.Printf("%v\n", gammaInt*epsilonInt)

}

func filterArray(inputArray []string, findingOx bool, filterIndex int) []string {
	count := 0
	for _, singleLine := range inputArray {
		charArray := []rune(singleLine)
		if charArray[filterIndex] == '1' {
			count += 1
		}
	}

	numOnes := count
	numZeroes := len(inputArray) - count

	moreOnes := false
	if numOnes >= numZeroes {
		moreOnes = true
	}

	var returnArray []string


	fmt.Printf("InputArr: %v\n",inputArray)

	// I hate the way this question is written because I think it means the same thing
	//for both cases and the flow below backs that up but I can't be sure until I run it.
	for _, singleLine := range inputArray {
		charArray := []rune(singleLine)
		if findingOx {
			if moreOnes {
				if charArray[filterIndex] == '1' {
					returnArray = append(returnArray, singleLine)
				}
			} else {
				if charArray[filterIndex] == '0' {
					returnArray = append(returnArray, singleLine)
				}
			}
		} else {
			if moreOnes {
				if charArray[filterIndex] == '0' {
					returnArray = append(returnArray, singleLine)
				}
			} else {
				if charArray[filterIndex] == '1' {
					returnArray = append(returnArray, singleLine)
				}
			}
		}
	}
	fmt.Printf("%v\n", returnArray)
	return returnArray
}

func findCo2Rate(inputArray []string) int {
	//var filteredArray []string
	//filteredArray:=make([]string,len(inputArray))

	filteredArray:=append([]string(nil),inputArray...)


	currentIndex := 0

	for len(filteredArray) > 1 {
		filteredArray = filterArray(filteredArray, false, currentIndex)
		currentIndex++
	}
	fmt.Printf("%v\n", filteredArray)
	theInt , _:= strconv.ParseInt(filteredArray[0],2,32)
	return int(theInt)

}

func findOxRate(inputArray []string) int {
	//var filteredArray []string
	//filteredArray:=make([]string,len(inputArray))

	filteredArray:=append([]string(nil),inputArray...)


	currentIndex := 0

	for len(filteredArray) > 1 {
		filteredArray = filterArray(filteredArray, true, currentIndex)
		currentIndex++
	}
	fmt.Printf("%v\n", filteredArray)
	theInt , _:= strconv.ParseInt(filteredArray[0],2,32)
	return int(theInt)

}

func solvePt2(inputLines []string) {
	oxRate := findOxRate(inputLines)
	co2Rate := findCo2Rate(inputLines)
	fmt.Printf("%v\n", oxRate*co2Rate)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
