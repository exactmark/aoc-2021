package Day09

import (
	"fmt"
	"strconv"
)

type cave struct {
	maxX     int
	maxY     int
	floorMap [][]floorPoint
}

type floorPoint struct {
	height int
	parent *cave
}

type coord struct {
	x int
	y int
}

func (c *cave) toString() string {
	returnStr := ""
	returnStr += fmt.Sprintf("Max x: %v\n", c.maxX)
	returnStr += fmt.Sprintf("Max y: %v\n", c.maxY)
	for _, row := range c.floorMap {
		for _, colVal := range row {
			returnStr += fmt.Sprintf("%v", colVal)
		}
		returnStr += "\n"
	}

	return returnStr
}

func (c *cave) populateFloor(lines []string) {
	c.maxX = len(lines[0])
	c.maxY = len(lines)
	c.floorMap = make([][]floorPoint, c.maxY)
	for y, singleLine := range lines {
		lineRunes := []rune(singleLine)
		thisRow := make([]floorPoint, len(singleLine))
		for x, val := range lineRunes {
			rowVal, _ := strconv.Atoi(string(val))
			thisRow[x] = floorPoint{
				height: rowVal,
				parent: c}
		}
		c.floorMap[y] = thisRow
	}

}

func (c *cave) getLowPoints() []coord {
	lowPointList := make([]coord, 0)
	for y := 0; y < c.maxY; y++ {
		for x := 0; x < c.maxX; x++ {
			thisCoord := coord{
				x: x,
				y: y,
			}
			if c.isLowPoint(thisCoord) {
				lowPointList = append(lowPointList, thisCoord)
			}
		}
	}
	return lowPointList
}

func (c *cave) isLowPoint(thisCoord coord) bool {
	thisHeight := c.floorMap[thisCoord.y][thisCoord.x].height
	neighborList:=c.floorMap[thisCoord.y][thisCoord.x].getNeighbors()
	return true
}

func (p *floorPoint) getNeighbors() []*floorPoint {
neighbors:=make([]*floorPoint,0)



return neighbors

}

func solvePt1(inputLines []string) {
	theCave := cave{
		maxX:     0,
		maxY:     0,
		floorMap: nil,
	}

	theCave.populateFloor(inputLines)
	lowPoints := theCave.getLowPoints()
	fmt.Printf("%v\n", lowPoints)
	fmt.Printf("%v\n", theCave.toString())
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
