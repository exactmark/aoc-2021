package Day14

import (
	"fmt"
	"sort"
	"strings"
)

func solvePt1(inputLines []string) {

	polymerArray := inputLines[0]

	insertMap := make(map[string]string)

	for _, val := range inputLines[2:] {
		insArray := strings.Split(val, " -> ")
		insertMap[insArray[0]] = insArray[1]
	}

	for i := 0; i < 40; i++ {
		nextString := ""

		for j := 0; j < len(polymerArray)-1; j++ {
			nextString+=string( polymerArray[j]) + insertMap[polymerArray[j:j+2]]
			//fmt.Printf("%v\n",polymerArray[j:j+2])
			//fmt.Printf("%v\n",insertMap[polymerArray[j:j+2]])
		}
		nextString += polymerArray[len(polymerArray)-1 :]
		polymerArray = nextString
		fmt.Printf("Step %v:%v\n",i+1)
	}

	runeCount:=make(map[rune]int)

	for _,val:=range []rune(polymerArray){
		if _,ok:=runeCount[val];ok{
			runeCount[val]++
		}else{
			runeCount[val]=1
		}
	}

	intArray:=make([]int,0)
	for _,val:=range runeCount{
		intArray=append(intArray,val)
	}

	sort.Ints(intArray)

	//fmt.Printf("%v\n", polymerArray)
	fmt.Printf("%v\n", runeCount)
	fmt.Printf("%v\n",intArray[len(intArray)-1]-intArray[0])
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
