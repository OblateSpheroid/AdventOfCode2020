package main

import (
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

func parseFile(s string) []Line {
	f, err := os.Open(s)
	defer f.Close()
	checkError(err)

	a := []Line{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var l Line
		tmp := strings.Split(scanner.Text(), " ")
		nums := strings.Split(tmp[0], "-")
		l.lower, _ = strconv.Atoi(nums[0])
		l.upper, _ = strconv.Atoi(nums[1])
		l.letter = tmp[1][0]
		l.pw = tmp[2]
		a = append(a, l)
	}

	return a
}
