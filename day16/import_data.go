package main

import (
	"aoc2020/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal - %v\n", e)
		os.Exit(1)
	}
}

func parseTickets(s string) [][]int {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	sl := [][]int{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		si := []int{}
		line := strings.Split(scanner.Text(), ",")
		for _, i := range line {
			n, _ := strconv.Atoi(i)
			helpers.AppendInPlace(&si, n)
		}
		sl = append(sl, si)
	}
	return sl
}

func parseMyTicket(s string) []int {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	sl := []int{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, i := range line {
			n, _ := strconv.Atoi(i)
			helpers.AppendInPlace(&sl, n)
		}
	}
	return sl
}

func parseRules(s string) map[string][]int {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	m := make(map[string][]int)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		c := strings.Split(line, ": ")
		m[c[0]] = []int{}
		n := strings.Split(c[1], " or ")
		for _, v := range n {
			ns := strings.Split(v, "-")
			min, _ := strconv.Atoi(ns[0])
			max, _ := strconv.Atoi(ns[1])
			r := helpers.MakeSeq(min, max)
			m[c[0]] = append(m[c[0]], r...)
		}
	}

	return m
}
