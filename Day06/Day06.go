package Day06

import (
	"fmt"
	"strconv"
	"strings"
)

func solvePt1(inputLines []string) {

	currentDay := make([]int, 9)
	dayDelta := make([]int, 9)

	inputArr := strings.Split(inputLines[0], ",")
	for _, v := range inputArr {
		newNumStr, _ := strconv.Atoi(v)
		currentDay[newNumStr]++
	}

	for i := 0; i < 80; i++ {
		dayDelta[6] = currentDay[0] + currentDay[7]
		dayDelta[8] = currentDay[0]
		dayDelta[7] = currentDay[8]
		for j := 1; j <= 6; j++ {
			dayDelta[j-1] = currentDay[j]
		}
		for j, _ := range currentDay {
			currentDay[j] = dayDelta[j]
		}
		//fmt.Printf("%v\n",currentDay)
		sum := 0
		for _, v := range currentDay {
			sum += v
		}
		fmt.Printf("Day %v: %v\n", i+1, sum)
	}

	fmt.Printf("%v\n", currentDay)

}

func solvePt2(inputLines []string) {

	currentDay := make([]int, 9)
	dayDelta := make([]int, 9)

	inputArr := strings.Split(inputLines[0], ",")
	for _, v := range inputArr {
		newNumStr, _ := strconv.Atoi(v)
		currentDay[newNumStr]++
	}

	for i := 0; i < 256; i++ {
		dayDelta[6] = currentDay[0] + currentDay[7]
		dayDelta[8] = currentDay[0]
		dayDelta[7] = currentDay[8]
		for j := 1; j <= 6; j++ {
			dayDelta[j-1] = currentDay[j]
		}
		for j, _ := range currentDay {
			currentDay[j] = dayDelta[j]
		}
		//fmt.Printf("%v\n",currentDay)
		sum := 0
		for _, v := range currentDay {
			sum += v
		}
		fmt.Printf("Day %v: %v\n", i+1, sum)
	}

	fmt.Printf("%v\n", currentDay)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
