package main

import (
	"fmt"
)

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

type Passport map[string]string

func checkPassport(p Passport) bool {
	// just check that all required fields are present
	for _, v := range required {
		if p[v] == "" {
			return false
		}
	}
	return true
}

func checkPassport2(p Passport) bool {
	// see all_rules.go for logic on each rule
	// loops through each rule, all must be true to return true
	if !checkPassport(p) {
		return false
	}
	if !checkByr(p) {
		return false
	}
	if !checkIyr(p) {
		return false
	}
	if !checkEyr(p) {
		return false
	}
	if !checkHgt(p) {
		return false
	}
	if !checkHcl(p) {
		return false
	}
	if !checkEcl(p) {
		return false
	}
	if !checkPid(p) {
		return false
	}
	return true
}

func countValid(ps []Passport, f func(p Passport) bool) int {
	i := 0
	for _, p := range ps {
		if f(p) {
			i++
		}
	}
	return i
}

func main() {
	ps := parseFile("data.txt") // see import_data.go for this function
	fmt.Printf("Answer 1: %d\n", countValid(ps, checkPassport))
	fmt.Printf("Answer 2: %d\n", countValid(ps, checkPassport2))
}
