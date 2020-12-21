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

func transpose(a [][]rune) [][]rune {
	// transpose a "matrix"
	if len(a) == 0 {
		return a
	}
	t := [][]rune{}
	for i := 0; i < len(a[0]); i++ {
		tmp := make([]rune, len(a))
		for j := range a {
			tmp = append(tmp, a[j][i])
		}
		t = append(t, tmp)
	}
	return t
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

/* Part 2 */
func maskv2(m, v string) []rune {
	// apply mask using v2 rules
	n := []rune(v) // can't modify string in place
	for i := 0; i < len(m); i++ {
		if m[i] != '0' {
			n[i] = rune(m[i]) // will produce runes with 'X'
		}
	}
	return n
}

func toBase10v2(s []rune) []int {
	// convert to all possible integers
	all := [][]rune{}
	for i, v := range s {
		if v == 'X' {
			if i == 0 {
				all = append(all, []rune("01"))
				continue
			}
			zeros := strings.Repeat("0", len(all[i-1]))
			ones := strings.Repeat("1", len(all[i-1]))
			all = append(all, []rune(zeros+ones))
			for j := range all[:i] {
				all[j] = append(all[j], all[j]...) // go back and double previous
			}
		} else {
			if i == 0 {
				all = append(all, []rune{v})
				continue
			}
			new := strings.Repeat(string(v), len(all[i-1]))
			all = append(all, []rune(new))
		}
	}
	t := transpose(all)
	// for each possibility, produce an integer it represents
	n := []int{}
	for _, v := range t {
		n = append(n, toBase10(v))
	}
	return n
}

func processv2(cs []Card) map[int]int {
	// mask memory address
	mem := make(map[int]int)
	for _, c := range cs {
		for _, i := range c.inst {
			s := toBase2(i[0])
			m := maskv2(c.mask, s)
			ints := toBase10v2(m)
			for _, j := range ints {
				mem[j] = i[1]
			}
		}
	}
	return mem
}

func main() {
	test1 := parseFile("test1.txt")
	t := process(test1)
	fmt.Println(count(t) == 165)

	test2 := parseFile("test2.txt")
	t2 := processv2(test2)
	fmt.Println(count(t2) == 208)

	data := parseFile("data.txt")
	m := process(data)
	fmt.Printf("Answer 1: %d\n", count(m))

	m2 := processv2(data)
	fmt.Printf("Answer 2: %d\n", count(m2))
}
