package Day07

import (
	"fmt"
	"strconv"
	"strings"
)

func getAbsDiff(left, right int) int {
	if left > right {
		return left - right
	} else {
		return right - left
	}
}

func determineFuel(crabMap *map[int]int, low int) int {

	sum := 0
	for k, v := range *crabMap {
		sum += getAbsDiff(k, low) * v
	}
	return sum
}

func determineFuelPt2(crabMap *map[int]int, low int) int {

	sum := 0
	for k, v := range *crabMap {
		steps := getAbsDiff(k, low)
		//sum of sequential numbers is
		//(n / 2)(first number + last number)
		fuelUsedPer := int((float32(steps)/2) * float32(steps+1))
		sum += fuelUsedPer * v
	}
	return sum
}

func solvePt1(inputLines []string) {

	crabsMap := make(map[int]int)

	crabsStrings := strings.Split(inputLines[0], ",")

	low, _ := strconv.Atoi(crabsStrings[0])
	high, _ := strconv.Atoi(crabsStrings[0])

	for _, singleCrab := range crabsStrings {
		crabNumb, _ := strconv.Atoi(singleCrab)
		if _, ok := crabsMap[crabNumb]; ok {
			crabsMap[crabNumb]++
		} else {
			crabsMap[crabNumb] = 1
		}
		if crabNumb > high {
			high = crabNumb
		}
		if crabNumb < low {
			low = crabNumb
		}
	}

	lowestX := low
	lowestFuel := determineFuel(&crabsMap, low)

	for i := low + 1; i < high; i++ {
		thisFuel := determineFuel(&crabsMap, i)
		if thisFuel < lowestFuel {
			lowestFuel = thisFuel
			lowestX = i
		}
		fmt.Printf("%v uses %v fuel\n", i, thisFuel)

	}

	fmt.Printf("%v-%v\n", low, high)
	fmt.Printf("%v\n", crabsMap)
	fmt.Printf("%v=%v\n", getAbsDiff(16, 2), getAbsDiff(2, 16))
	fmt.Printf("%v uses %v fuel\n", lowestX, lowestFuel)
}

func solvePt2(inputLines []string) {
	crabsMap := make(map[int]int)

	crabsStrings := strings.Split(inputLines[0], ",")

	low, _ := strconv.Atoi(crabsStrings[0])
	high, _ := strconv.Atoi(crabsStrings[0])

	for _, singleCrab := range crabsStrings {
		crabNumb, _ := strconv.Atoi(singleCrab)
		if _, ok := crabsMap[crabNumb]; ok {
			crabsMap[crabNumb]++
		} else {
			crabsMap[crabNumb] = 1
		}
		if crabNumb > high {
			high = crabNumb
		}
		if crabNumb < low {
			low = crabNumb
		}
	}

	lowestX := low
	lowestFuel := determineFuelPt2(&crabsMap, low)

	for i := low + 1; i < high; i++ {
		thisFuel := determineFuelPt2(&crabsMap, i)
		if thisFuel < lowestFuel {
			lowestFuel = thisFuel
			lowestX = i
		}
		fmt.Printf("%v uses %v fuel\n", i, thisFuel)

	}

	fmt.Printf("%v-%v\n", low, high)
	fmt.Printf("%v\n", crabsMap)
	fmt.Printf("%v=%v\n", getAbsDiff(16, 2), getAbsDiff(2, 16))
	fmt.Printf("%v uses %v fuel\n", lowestX, lowestFuel)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
