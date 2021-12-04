package Day03

import (
	"fmt"
	"strconv"
)

func byteArrayToInt(bitArray []rune) int64 {
	intVal, _ := strconv.ParseInt(string(bitArray), 2, 64)
	return intVal
}

func solvePt1(inputLines []string) {
	countingArray := make([]int, len(inputLines[0]))
	for _, singleLine := range inputLines {
		charArray := []rune(singleLine)
		for k, v := range charArray {
			if v == '1' {
				countingArray[k] += 1
			}
		}

	}

	gammaBits := make([]rune, len(countingArray))
	epsilonBits := make([]rune, len(countingArray))
	for k, v := range countingArray {
		if float32(v)/float32(len(inputLines)) > 0.5 {
			gammaBits[k] = '1'
			epsilonBits[k] = '0'
		} else {
			gammaBits[k] = '0'
			epsilonBits[k] = '1'
		}
	}

	gammaInt := byteArrayToInt(gammaBits)
	epsilonInt := byteArrayToInt(epsilonBits)
	fmt.Printf("%v,%v\n", gammaBits, epsilonBits)
	fmt.Printf("%v,%v\n", gammaInt, epsilonInt)
	fmt.Printf("%v\n", gammaInt* epsilonInt)

}

func filterArray(inputArray[]string,keepHigh bool,filterIndex int)[]string{
	count:=0
	for _, singleLine := range inputArray {
		charArray := []rune(singleLine)
		if charArray[filterIndex]=='1'{
			count+=1
		}
	}

	signifier:='0'
	if count>(len(inputArray)-count){
		signifier='1'
	}

	var returnArray []string



	for _, singleLine := range inputArray {
		charArray := []rune(singleLine)
		if charArray[filterIndex]==signifier && keepHigh{
		 returnArray=	append(returnArray, singleLine)
		}
	}
	fmt.Printf("%v\n",returnArray)
	return returnArray
}

func findOxRate(inputArray[]string)int{
	var filteredArray []string
	//filteredArray:=make([]string,len(inputArray))

	copy(filteredArray,inputArray)

	currentIndex :=0
	fmt.Printf("%v\n",filteredArray)

	for len(filteredArray)>1{
		filteredArray = filterArray(filteredArray,true, currentIndex)
	}
	fmt.Printf("%v\n",filteredArray)
	return 0
}

func solvePt2(inputLines []string) {
	oxRate:=findOxRate(inputLines)
	fmt.Printf("%v\n",oxRate)
}

func Solve(inputLines []string) {
	//solvePt1(inputLines)
	solvePt2(inputLines)
}
