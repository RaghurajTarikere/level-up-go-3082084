package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	var myStack stack

	for _, bracket := range expr {
		if isOpeningBracketType(bracket) {
			myStack.push(bracket)
		} else if isClosingBracketType(bracket) {
			curBracket := myStack.pop()
			if curBracket != getMatchingBracketType(bracket) {
				return false
			}
		}
	}
	return true
}

func isOpeningBracketType(e rune) bool {
	return e == '{' || e == '[' || e == '('
}

func isClosingBracketType(e rune) bool {
	return e == '}' || e == ']' || e == ')'
}

func getMatchingBracketType(e rune) rune {
	switch e {
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	default:
		log.Fatal("Invalid bracket type")
	}
	return ' '
}

type stack struct {
	elems []rune
}

func (s *stack) push(e rune) {
	s.elems = append(s.elems, e)
}

func (s *stack) pop() rune {
	if len(s.elems) == 0 {
		log.Fatal("No elems in stack to pop")
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
