package Day17

import "fmt"

type probe struct {
	xPos int
	yPos int
	xVel int
	yVel int
	yMax int
}

type targetSpace struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func (p *probe) processTime() {
	p.xPos += p.xVel
	p.yPos += p.yVel
	if p.yMax < p.yPos {
		p.yMax = p.yPos
	}

	if p.xVel > 0 {
		p.xVel -= 1
	} else if p.xVel < 0 {
		p.xVel += 1
	}
	p.yVel -= 1
}

func (p *probe) isInTarget(t targetSpace) bool {
	if p.xPos >= t.xMin && p.xPos <= t.xMax && p.yPos >= t.yMin && p.yPos <= t.yMax {
		return true
	}
	return false
}

func (p *probe) couldStillReach(t targetSpace) bool {
	if p.yPos < t.yMin {
		return false
	}
	if p.xPos < t.xMin && p.xVel < 1 {
		return false
	}
	if p.xPos > t.xMax && p.xVel < (-1) {
		return false
	}
	return true
}

func solvePt1(inputLines []string) {

	t := targetSpace{
		xMin: 70,
		xMax: 125,
		yMin: -159,
		yMax: -121,
	}

	//t := targetSpace{
	//	xMin: 20,
	//	xMax: 30,
	//	yMin: -10,
	//	yMax: -5,
	//}

	maxVel := 1200

	bestY := 0

	for x := 0; x < maxVel; x++ {
		for y := 0; y < maxVel; y++ {
			//fmt.Printf("Trying %v,%v\n", x, y)
			p := probe{
				xPos: 0,
				yPos: 0,
				xVel: x,
				yVel: y,
				yMax: 0,
			}
			for !p.isInTarget(t) && p.couldStillReach(t) {
				p.processTime()
			}
			if p.isInTarget(t) {
				fmt.Printf("Success!\n")
				fmt.Printf("%v,%v,%v\n", x, y, p.yMax)
				if p.yMax > bestY {
					bestY = p.yMax
				}
			}

		}
	}
	fmt.Printf("Best: %v\n", bestY)
}

func solvePt2(inputLines []string) {

	t := targetSpace{
		xMin: 70,
		xMax: 125,
		yMin: -159,
		yMax: -121,
	}

	//t := targetSpace{
	//	xMin: 20,
	//	xMax: 30,
	//	yMin: -10,
	//	yMax: -5,
	//}

	maxVel := 1200

	bestY := 0
numSol:=0

	for x := 0; x < maxVel; x++ {
		for y := -160; y < maxVel; y++ {
			//fmt.Printf("Trying %v,%v\n", x, y)
			p := probe{
				xPos: 0,
				yPos: 0,
				xVel: x,
				yVel: y,
				yMax: 0,
			}
			for !p.isInTarget(t) && p.couldStillReach(t) {
				p.processTime()
			}
			if p.isInTarget(t) {
				fmt.Printf("Success!\n")
				fmt.Printf("%v,%v,%v\n", x, y, p.yMax)
				if p.yMax > bestY {
					bestY = p.yMax
				}
				numSol++
			}

		}
	}
	fmt.Printf("Best: %v\n", bestY)
	fmt.Printf("Found: %v\n", numSol)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
