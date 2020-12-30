package main

import (
	"fmt"
	"strconv"
	"strings"
)

type P struct {
	result int
	str    []string
}

func evaluate(s []string, starting int, ops string, ch chan P) int {
	// recursive function to walk through string, evaluating left to right
	if s[0] == "+" || s[0] == "*" {
		// apply with new operator
		return evaluate(s[1:], starting, s[0], ch)
	} else if s[0] == "(" {
		// create new goroutine for ()
		ch2 := make(chan P)
		go evaluate(s[1:], 0, "+", ch2) // start over within ()
		v := <-ch2
		if ops == "+" {
			return evaluate(v.str, starting+v.result, ops, ch)
		} else {
			return evaluate(v.str, starting*v.result, ops, ch)
		}
	} else if s[0] == ")" {
		// end goroutine and return result
		if len(s) > 1 {
			ch <- P{result: starting, str: s[1:]}
		} else {
			ch <- P{result: starting, str: []string{""}}
		}
	} else if s[0] == "" {
		// got END signal, return result so far
		return starting
	} else {
		// character is a number, add or multiply based on operator
		n, _ := strconv.Atoi(s[0])
		if len(s) > 1 {
			if ops == "+" {
				return evaluate(s[1:], n+starting, ops, ch)
			} else if ops == "*" {
				return evaluate(s[1:], n*starting, ops, ch)
			}
		} else { // reached end, final calculation
			if ops == "+" {
				return n + starting
			} else if ops == "*" {
				return n * starting
			}
		}
	}
	return -1
}

func evaluateData(data [][]string) int {
	c := 0
	ch := make(chan P)
	for _, line := range data {
		c += evaluate(line, 0, "+", ch)
	}
	return c
}

func main() {
	test := strings.Split(`((2+4*9)*(6+9*8+6)+6)+2+4*2`, "")
	test2 := strings.Split(`5*9*(7*3*3+9*3+(8+6*4))`, "")
	fmt.Println(evaluate(test, 0, "+", make(chan P)) == 13632)
	fmt.Println(evaluate(test2, 0, "+", make(chan P)) == 12240)

	data := parseFile("data.txt")
	fmt.Printf("Answer 1: %d\n", evaluateData(data))
}
