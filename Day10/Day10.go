package Day10

import (
	"fmt"
	"strings"
)

type stack struct {
	runes []rune
}

func (s *stack) append(newRune rune) {
	s.runes = append(s.runes, newRune)
}

func (s *stack) pop() rune {
	retRune := s.runes[len(s.runes)-1]
	s.runes = s.runes[:len(s.runes)-1]
	return retRune
}

func isMatch(opener, closer rune) bool {
	switch opener {
	case '(':
		if closer == ')' {
			return true
		}
	case '[':
		if closer == ']' {
			return true
		}
	case '{':
		if closer == '}' {
			return true
		}
	case '<':
		if closer == '>' {
			return true
		}
	}
	return false
}

func findMismatch(singleLine string) rune {
	theRunes := []rune(singleLine)
	theStack := stack{runes: make([]rune, 0)}
	for _, val := range theRunes {
		if strings.ContainsAny(string(val), "([{<") {
			theStack.append(val)
		} else {
			theMatch := theStack.pop()
			if isMatch(theMatch, val) {
				//	hooray!
			} else {
				return val
			}
		}
		//fmt.Printf("%v\n", theStack)
	}
	return '.'
}

func solvePt1(inputLines []string) {

	sum:=0

	for _, singleLine := range inputLines {
		mismatch := findMismatch(singleLine)
		fmt.Printf("%v\n", string(mismatch))
		switch mismatch {
		case ')':
			sum+=3
		case ']':
			sum+=57
		case '}':
			sum+=1197
		case '>':
			sum+=25137
		}
	}
	fmt.Printf("%v\n",sum)

}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
