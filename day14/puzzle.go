package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	mask string
	inst [][2]int // memory, value pairs
}

func toBase2(i int) string {
	// string representing integer in base 2
	s := strconv.FormatInt(int64(i), 2)
	return strings.Repeat("0", 36-len(s)) + s // left pad w 0s
}

func toBase10(s []rune) int {
	// convert back to a normal integer
	n := 0
	for i := 0; i < len(s); i++ {
		e := float64(len(s) - (i + 1))
		y, _ := strconv.Atoi(string(s[i]))
		if y > 0 {
			n += int(math.Pow(2, e))
		}
	}
	return n
}

func mask(m, v string) []rune {
	// apply mask
	n := []rune(v) // can't modify string in place
	for i := 0; i < len(m); i++ {
		if m[i] != 'X' {
			n[i] = rune(m[i])
		}
	}
	return n
}

func process(cs []Card) map[int]int {
	mem := make(map[int]int)
	for _, c := range cs {
		for _, i := range c.inst {
			s := toBase2(i[1])
			m := mask(c.mask, s)
			mem[i[0]] = toBase10(m)
		}
	}
	return mem
}

func count(m map[int]int) int {
	c := 0
	for _, v := range m {
		c += v
	}
	return c
}

func main() {
	data := parseFile("data.txt")
	m := process(data)
	fmt.Println(count(m))
}
