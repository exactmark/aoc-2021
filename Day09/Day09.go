package Day09

import (
	"fmt"
	"sort"
	"strconv"
)

type cave struct {
	maxX     int
	maxY     int
	floorMap map[coord]floorPoint
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
	for y := 0; y < c.maxY; y++ {
		for x := 0; x < c.maxX; x++ {
			val, _ := c.floorMap[coord{
				x: x,
				y: y,
			}]
			returnStr += fmt.Sprintf("%v", val.height)
		}
		returnStr += "\n"
	}

	return returnStr
}

func (c *cave) populateFloor(lines []string) {
	c.maxX = len(lines[0])
	c.maxY = len(lines)
	c.floorMap = make(map[coord]floorPoint)
	for y, singleLine := range lines {
		lineRunes := []rune(singleLine)
		for x, val := range lineRunes {
			rowVal, _ := strconv.Atoi(string(val))
			c.floorMap[coord{
				x: x,
				y: y,
			}] = floorPoint{
				height: rowVal,
				parent: c,
			}
		}
	}

}

func (c *cave) getLowPoints() []coord {
	lowPointList := make([]coord, 0)
	for k, _ := range c.floorMap {
		if c.isLowPoint(k) {
			lowPointList = append(lowPointList, k)
		}
	}
	return lowPointList
}

func (c *cave) isLowPoint(thisCoord coord) bool {
	sutHeight := c.floorMap[thisCoord].height
	testCoord := thisCoord
	testCoord.x += -1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height <= sutHeight {
			return false
		}
	}
	testCoord = thisCoord
	testCoord.x += 1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height <= sutHeight {
			return false
		}
	}
	testCoord = thisCoord
	testCoord.y += -1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height <= sutHeight {
			return false
		}
	}
	testCoord = thisCoord
	testCoord.y += 1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height <= sutHeight {
			return false
		}
	}
	return true
}

func (c *cave) getBasinNeighbors(thisCoord coord) []coord {
	neighbors := make([]coord, 0)
	testCoord := thisCoord
	testCoord.x += -1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height != 9 {
			neighbors = append(neighbors, testCoord)
		}
	}
	testCoord = thisCoord
	testCoord.x += 1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height != 9 {
			neighbors = append(neighbors, testCoord)
		}
	}
	testCoord = thisCoord
	testCoord.y += -1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height != 9 {
			neighbors = append(neighbors, testCoord)
		}
	}
	testCoord = thisCoord
	testCoord.y += 1
	if val, ok := c.floorMap[testCoord]; ok {
		if val.height != 9 {
			neighbors = append(neighbors, testCoord)
		}
	}
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
	sum := 0
	for _, val := range lowPoints {
		fmt.Printf("%v\n", theCave.floorMap[val].height+1)
		sum += theCave.floorMap[val].height + 1
	}
	fmt.Printf("sum= %v\n", sum)
}

func solvePt2(inputLines []string) {
	theCave := cave{
		maxX:     0,
		maxY:     0,
		floorMap: nil,
	}

	theCave.populateFloor(inputLines)
	lowPoints := theCave.getLowPoints()
	fmt.Printf("%v\n", lowPoints)
	fmt.Printf("%v\n", theCave.toString())
	sum := 0
	for _, val := range lowPoints {
		fmt.Printf("%v\n", theCave.floorMap[val].height+1)
		sum += theCave.floorMap[val].height + 1
	}

	basinSizes:=make([]int,0)

	for _,val :=range lowPoints{
		thisBasin:=make(map[coord]bool)
		thisBasin[val]=true
		lastSize:=0
		for len(thisBasin)>lastSize{
			lastSize=len(thisBasin)
			for checkPoint,_:=range thisBasin{
				for _,newPoint:=range theCave.getBasinNeighbors(checkPoint){
					thisBasin[newPoint]=true
				}
			}

		}
		basinSizes=append(basinSizes,len(thisBasin))
	}

	sort.Ints(basinSizes)
	basinSizes=basinSizes[len(basinSizes)-3:]



	fmt.Printf("basins= %v\n", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
