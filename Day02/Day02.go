package Day02

import (
	"fmt"
	"strconv"
	"strings"
)

type submarine struct {
	xDist int
	depth int
	aim int
}

func (v *submarine) takeCommand(direction string, distance int) {

	if direction[0] == 'f' {
		v.xDist += distance
	} else if direction[0] == 'd' {
		v.depth += distance
	} else {
		v.depth -= distance
	}
	if v.depth<0{
		fmt.Printf("Breach!\n")
	}
}

func solvePt1(inputLines []string) {
	mySub := submarine{
		xDist: 0,
		depth: 0,
	}

	for _, singleLine := range inputLines {
		lineSplit := strings.Split(singleLine, " ")
		distance, _ := strconv.Atoi(lineSplit[1])
		direction := lineSplit[0]
		mySub.takeCommand(direction, distance)
	}

	fmt.Printf("%v*%v=%v\n", mySub.xDist, mySub.depth, mySub.xDist*mySub.depth)

}


func (v *submarine) takeAimedCommand(direction string, distance int) {

	if direction[0] == 'f' {
		v.xDist += distance
		v.depth += distance * v.aim
	} else if direction[0] == 'd' {
		v.aim += distance
	} else {
		v.aim -= distance
	}
	if v.depth<0{
		fmt.Printf("Breach!\n")
	}
}


func solvePt2(inputLines []string) {
	mySub := submarine{
		xDist: 0,
		depth: 0,
		aim:0,
	}

	for _, singleLine := range inputLines {
		lineSplit := strings.Split(singleLine, " ")
		distance, _ := strconv.Atoi(lineSplit[1])
		direction := lineSplit[0]
		mySub.takeAimedCommand(direction, distance)
	}

	fmt.Printf("%v*%v=%v\n", mySub.xDist, mySub.depth, mySub.xDist*mySub.depth)

}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
