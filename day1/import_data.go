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

	a := make([]int, 0, 200)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		tmp, _ := strconv.Atoi(scanner.Text())
		a = append(a, tmp)
	}

	return a
}
