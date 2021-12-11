package Day11

import (
	"fmt"
	"strconv"
)

type coord struct {
	x int
	y int
}

type floorMap struct {
	octMap     map[coord]*octopus
	maxX       int
	maxY       int
	numFlashes int
}

func (m *floorMap) populate(lines []string) {

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			energy, _ := strconv.Atoi(string(lines[y][x]))
			m.octMap[coord{
				x: x,
				y: y,
			}] = &octopus{
				energy:      energy,
				floorParent: m,
				neighbors:   make([]*octopus, 0),
				hasFlashed:  false,
			}
			if x > m.maxX {
				m.maxX = x
			}
		}
		if y > m.maxY {
			m.maxY = y
		}
	}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			neighborList := make([]*octopus, 0)
			for yN := -1; yN <= 1; yN++ {
				for xN := -1; xN <= 1; xN++ {
					if xN != 0 || yN != 0 {
						if _, ok := m.octMap[coord{
							x: x + xN,
							y: y + yN,
						}]; ok {
							neighborList = append(neighborList, m.octMap[coord{
								x: x + xN,
								y: y + yN,
							}])
						}
					}
				}
			}
			m.octMap[coord{
				x: x,
				y: y,
			}].neighbors = neighborList
		}
	}
}

func (m *floorMap) print() {
	thePrint := ""
	for y := 0; y <= m.maxY; y++ {
		for x := 0; x <= m.maxX; x++ {
			thePrint += fmt.Sprintf("%v", m.octMap[coord{
				x: x,
				y: y,
			}].energy)
		}
		thePrint += "\n"
	}
	fmt.Printf(thePrint)
}

func (m *floorMap) processStep() {
	m.addBaseEnergy()
	m.processFlashes()
}

func (m *floorMap) addBaseEnergy() {
	for _, val := range m.octMap {
		val.energy += 1
		val.hasFlashed = false
	}
}

func (m *floorMap) processFlashes() {
	for _, val := range m.octMap {
		val.maybeFlash()
	}
}

func (m *floorMap) allFlashed()bool {
	for _,val:=range m.octMap{
		if !val.hasFlashed{
			return false
		}
	}
	return true
}

type octopus struct {
	energy      int
	floorParent *floorMap
	neighbors   []*octopus
	hasFlashed  bool
}

func (o *octopus) maybeFlash() {
	if o.hasFlashed || o.energy <= 9 {
		return
	}
	o.floorParent.numFlashes += 1
	o.energy = 0
	o.hasFlashed = true
	for _, val := range o.neighbors {
		if !val.hasFlashed {
			val.energy += 1
			val.maybeFlash()
		}
	}

}

func solvePt1(inputLines []string) {
	theMap := floorMap{octMap: make(map[coord]*octopus)}
	theMap.populate(inputLines)
	fmt.Printf("%v\n", theMap)
	fmt.Printf("%v\n", theMap.octMap[coord{
		x: 2,
		y: 2,
	}])
	fmt.Printf("%v\n", theMap.octMap[coord{
		x: 0,
		y: 0,
	}])
	theMap.print()
	for i := 1; i <= 100; i++ {
		theMap.processStep()
		//fmt.Printf("\nAfter step %v:\n", i)
		//theMap.print()
	}
	fmt.Printf("numFlashes: %v\n",theMap.numFlashes)
}

func solvePt2(inputLines []string) {
	theMap := floorMap{octMap: make(map[coord]*octopus)}
	theMap.populate(inputLines)
	//fmt.Printf("%v\n", theMap)
	//fmt.Printf("%v\n", theMap.octMap[coord{
	//	x: 2,
	//	y: 2,
	//}])
	//fmt.Printf("%v\n", theMap.octMap[coord{
	//	x: 0,
	//	y: 0,
	//}])
	//theMap.print()
	for i := 1; i <= 3000; i++ {
		theMap.processStep()
		if theMap.allFlashed(){
			fmt.Printf("allFlashed on %v\n",i)
			break
		}
		//fmt.Printf("\nAfter step %v:\n", i)
		//theMap.print()
	}
	fmt.Printf("numFlashes: %v\n",theMap.numFlashes)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
