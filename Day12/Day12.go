package Day12

import (
	"fmt"
	"strings"
)

type nodeMap struct {
	nMap          map[string]node
	solutionArray [][]string
}

type node struct {
	neighbors map[string]bool
	isBig     bool
	numVisits int
}

func solvePt1(inputLines []string) {
	theMap := nodeMap{nMap: make(map[string]node)}
	theMap.populateNodeMap(inputLines)

	//fmt.Printf("%v\n", theMap)
	//for k,val:=range theMap.nMap{
	//	fmt.Printf("%v:%v\n", k,val)
	//}

	theMap.solutionArray = make([][]string, 0)
	partialPath := make([]string, 0)

	theMap.maybeAddNextNode(partialPath, "start")

	fmt.Printf("%v\n", theMap.solutionArray)
	fmt.Printf("%v\n", len(theMap.solutionArray))
}

func (n *nodeMap) populateNodeMap(lines []string) {
	for _, singleLine := range lines {
		splits := strings.Split(singleLine, "-")
		n.addNode(splits)
	}
}

func (n *nodeMap) addNode(splits []string) {
	src := splits[0]
	dest := splits[1]
	n.addNodeHelper(src, dest)
	n.addNodeHelper(dest, src)
}

func (n *nodeMap) addNodeHelper(src, dest string) {
	if val, ok := n.nMap[src]; ok {
		val.neighbors[dest] = true
	} else {
		srcIsRevisitable := false
		if strings.ToUpper(src) == src {
			srcIsRevisitable = true
		}
		n.nMap[src] = node{
			neighbors: make(map[string]bool),
			isBig:     srcIsRevisitable,
		}
		n.nMap[src].neighbors[dest] = true
	}
}

func (n *nodeMap) maybeAddNextNode(path []string, next string) {
	if next == "end" {
		path = append(path, next)
		n.solutionArray = append(n.solutionArray, path)
	}
	//	if next.isBig or next is not in path
	//	 add next to path
	//	 for each neighbor
	//		make copy of path
	//  	send copy and next to maybeAddNextNode
	if n.nMap[next].isBig || !contains(path, next) {
		path = append(path, next)
		for key, _ := range n.nMap[next].neighbors {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			n.maybeAddNextNode(pathCopy, key)
		}
	}
}

func contains(slice []string, target string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
