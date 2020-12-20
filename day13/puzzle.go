package main

import (
	"fmt"
	"math/big"
)

func leastWait(t int, ids []int) (int, int) {
	wait := t // start with high min wait time
	bus := -1
	for _, id := range ids {
		if id > 0 {
			tmp := (id - (t % id))
			if tmp < wait {
				wait = tmp
				bus = id
			}
		}
	}
	return bus, wait
}

// Chinese remainder theorem
// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

// convert bus IDs to a, n for CRT
func process(data []int) ([]*big.Int, []*big.Int) {
	n := []*big.Int{}
	a := []*big.Int{}
	for i, v := range data {
		if v > 0 {
			a = append(a, big.NewInt(int64(v-i)))
			n = append(n, big.NewInt(int64(v)))
		}
	}
	a[0] = big.NewInt(int64(0))
	return a, n
}

func main() {
	t, b := parseFile("test1.txt")
	bs, w := leastWait(t, b)
	fmt.Println((bs * w) == 295)
	test, _ := crt(process(b))
	fmt.Println(fmt.Sprint(test) == "1068781")

	time, buses := parseFile("data.txt")
	bus, wait := leastWait(time, buses)
	fmt.Printf("Answer 1: %d\n", bus*wait)
	sol, _ := crt(process(buses))
	fmt.Printf("Answer 2: %d\n", sol)
}
