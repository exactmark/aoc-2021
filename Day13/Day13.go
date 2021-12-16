package Day13

import (
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type transparentPaper struct {
	grid map[coord]bool
	maxX int
	maxY int
}

func (p *transparentPaper) addNode(line string) {
	stringSplit := strings.Split(line, ",")
	x, _ := strconv.Atoi(stringSplit[0])
	y, _ := strconv.Atoi(stringSplit[1])
	p.grid[coord{
		x: x,
		y: y,
	}] = true
	if x > p.maxX {
		p.maxX = x
	}
	if y > p.maxY {
		p.maxY = y
	}
}

func (p *transparentPaper) printPaper() {
	for y := 0; y <= p.maxY; y++ {
		for x := 0; x <= p.maxX; x++ {
			if _, ok := p.grid[coord{
				x: x,
				y: y,
			}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (p *transparentPaper) doFold(line string) {
	stringSplit := strings.Split(line, " ")
	stringSplit = strings.Split(stringSplit[2], "=")
	amount, _ := strconv.Atoi(stringSplit[1])
	if stringSplit[0] == "x" {
		p.doXSplit(amount)
	} else {
		p.doYSplit(amount)
	}
}

func (p *transparentPaper) doXSplit(amount int) {
	for key,_:=range p.grid{
		if key.x>amount{
			diff := key.x-amount
			delete(p.grid,key)
			p.grid[coord{
				x: key.x-2*diff,
				y: key.y,
			}]=true
		}
	}
	p.maxX=amount-1
}

func (p *transparentPaper) doYSplit(amount int) {
	for key,_:=range p.grid{
		if key.y>amount{
			diff := key.y-amount
			delete(p.grid,key)
			p.grid[coord{
				x: key.x,
				y: key.y-2*diff,
			}]=true
		}
	}
	p.maxY=amount-1
}

func solvePt1(inputLines []string) {

	paper := transparentPaper{grid: map[coord]bool{}}

	for _, singleLine := range inputLines {
		if strings.Contains(singleLine, ",") {
			paper.addNode(singleLine)
		}
	}
	for _, singleLine := range inputLines {
		if strings.Contains(singleLine, "fold") {
			paper.doFold(singleLine)
			break
		}
		//fmt.Printf("\n")
		//paper.printPaper()
	}

	paper.printPaper()
fmt.Printf("%v\n",len(paper.grid))
}

func solvePt2(inputLines []string) {
	paper := transparentPaper{grid: map[coord]bool{}}

	for _, singleLine := range inputLines {
		if strings.Contains(singleLine, ",") {
			paper.addNode(singleLine)
		}
	}
	for _, singleLine := range inputLines {
		if strings.Contains(singleLine, "fold") {
			paper.doFold(singleLine)
			//break
		}
		//fmt.Printf("\n")
		//paper.printPaper()
	}

	paper.printPaper()

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
