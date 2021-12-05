package Day05

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

type oceanMap struct {
	knownSpaces map[coord]int
	largestX    int
	largestY    int
}

func numStringArrToCoord(inputStrArr []string) coord {
	xPart, _ := strconv.Atoi(inputStrArr[0])
	yPart, _ := strconv.Atoi(inputStrArr[1])
	return coord{
		x: xPart,
		y: yPart,
	}

}

func lineToCoords(inputString string) (coord, coord) {

	pieces := strings.Split(inputString, " -> ")

	startString := strings.Split(pieces[0], ",")
	endString := strings.Split(pieces[1], ",")
	return numStringArrToCoord(startString), numStringArrToCoord(endString)

}

func (b *oceanMap) toString() string {
	returnString := ""
	for y := 0; y < b.largestY; y++ {
		for x := 0; x < b.largestX; x++ {
			thisCoord := coord{
				x: x,
				y: y,
			}
			if val, ok := b.knownSpaces[thisCoord]; ok {
				returnString += fmt.Sprintf("%v", val)
			} else {
				returnString += "."
			}
		}
		returnString += "\n"
	}
	//returnString += fmt.Sprintf("remaining: %v\n", len(b.heldNumbers))
	//returnString += fmt.Sprintf("held: %v\n", b.heldNumbers)
	returnString += fmt.Sprintf("extent: %v,%v\n", b.largestX, b.largestY)
	returnString += fmt.Sprintf("occupied: %v\n", b.knownSpaces)
	return returnString
}

func (b *oceanMap) addLine(startPt coord, endPt coord) {
	if startPt.x == endPt.x {
		//figure out which is bigger
		lowY := startPt.y
		highY := endPt.y
		if startPt.y > endPt.y {
			highY = startPt.y
			lowY = endPt.y
		}
		// add those points
		for i := lowY; i <= highY; i++ {
			b.incrementPt(startPt.x, i)
		}
	} else if startPt.y == endPt.y {
		lowX := startPt.x
		highX := endPt.x
		if startPt.x > endPt.x {
			highX = startPt.x
			lowX = endPt.x
		}
		for i := lowX; i <= highX; i++ {
			b.incrementPt(i, startPt.y)
		}
	}
}

func (b *oceanMap) addLineWithDiagonals(startPt coord, endPt coord) {
	if startPt.x == endPt.x || startPt.y == endPt.y {
		b.addLine(startPt, endPt)
	} else {
		xInc := 1
		yInc := 1
		if startPt.x > endPt.x {
			xInc = -1
		}
		if startPt.y > endPt.y {
			yInc = -1
		}
		currentX := startPt.x
		currentY := startPt.y
		b.incrementPt(currentX, currentY)
		for true {
			currentX += xInc
			currentY += yInc
			b.incrementPt(currentX, currentY)
			if currentX == endPt.x {
				break
			}
		}
	}
}

func (b *oceanMap) incrementPt(x int, y int) {
	thisCoord := coord{
		x: x,
		y: y,
	}
	if _, ok := b.knownSpaces[thisCoord]; ok {
		b.knownSpaces[thisCoord]++
	} else {
		b.knownSpaces[thisCoord] = 1
	}
	if x > b.largestX {
		b.largestX = x
	}
	if y > b.largestY {
		b.largestY = y
	}
}

func solvePt1(inputLines []string) {

	thisMap := oceanMap{
		knownSpaces: make(map[coord]int),
		largestX:    0,
		largestY:    0,
	}

	for _, singleLine := range inputLines {
		startPoint, endPoint := lineToCoords(singleLine)
		thisMap.addLine(startPoint, endPoint)
	}

	fmt.Printf("%v\n", thisMap.toString())

	//	count the high spots
	sum := 0

	for _, val := range thisMap.knownSpaces {
		if val > 1 {
			sum++
		}
	}

	fmt.Printf("High: %v\n", sum)

}

func solvePt2(inputLines []string) {

	thisMap := oceanMap{
		knownSpaces: make(map[coord]int),
		largestX:    0,
		largestY:    0,
	}

	for _, singleLine := range inputLines {
		startPoint, endPoint := lineToCoords(singleLine)
		thisMap.addLineWithDiagonals(startPoint, endPoint)
		//fmt.Printf("%v,%v\n", startPoint, endPoint)
	}

	//fmt.Printf("%v\n", thisMap.toString())

	//	count the high spots
	sum := 0

	for _, val := range thisMap.knownSpaces {
		if val > 1 {
			sum++
		}
	}

	fmt.Printf("High: %v\n", sum)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
