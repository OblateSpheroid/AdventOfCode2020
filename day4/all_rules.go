package main

import (
	"strconv"
)

func checkByr(p Passport) bool {
	if len(p["byr"]) != 4 {
		return false
	}
	test, err := strconv.Atoi(p["byr"])
	if err != nil {
		return false
	}
	if test < 1920 || test > 2002 {
		return false
	}
	return true
}

func checkIyr(p Passport) bool {
	if len(p["iyr"]) != 4 {
		return false
	}
	test, err := strconv.Atoi(p["iyr"])
	if err != nil {
		return false
	}
	if test < 2010 || test > 2020 {
		return false
	}
	return true
}

func checkEyr(p Passport) bool {
	if len(p["eyr"]) != 4 {
		return false
	}
	test, err := strconv.Atoi(p["eyr"])
	if err != nil {
		return false
	}
	if test < 2020 || test > 2030 {
		return false
	}
	return true
}

func checkHgt(p Passport) bool {
	hgt := p["hgt"]
	hgt_ln := len(hgt)
	test, err := strconv.Atoi(hgt[:hgt_ln-2])
	if err != nil {
		return false
	}
	hgt_sfx := hgt[hgt_ln-2 : hgt_ln]
	if hgt_sfx == "cm" {
		if test < 150 || test > 193 {
			return false
		}
	} else if hgt_sfx == "in" {
		if test < 59 || test > 76 {
			return false
		}
	} else {
		return false
	}
	return true
}

func checkHcl(p Passport) bool {
	if p["hcl"][0] != 35 {
		return false
	}
	for _, v := range p["hcl"][1:] {
		if (v < 48 || v > 57) && (v < 97 || v > 122) { //ranges for alphanumeric
			return false
		}
	}
	return true
}

func checkEcl(p Passport) bool {
	allowed_ecl := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	good_ecl := false
	for _, v := range allowed_ecl {
		if p["ecl"] == v {
			good_ecl = true
		}
	}
	if !good_ecl {
		return false
	}
	return true
}

func checkPid(p Passport) bool {
	if len(p["pid"]) != 9 {
		return false
	}
	_, err := strconv.Atoi(p["pid"])
	if err != nil {
		return false
	}
	return true
}
