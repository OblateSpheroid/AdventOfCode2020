package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal - %v\n", e)
		os.Exit(1)
	}
}

func parseFile(s string) []int {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	sl := []int{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		sl = append(sl, i)
	}

	return sl
}
