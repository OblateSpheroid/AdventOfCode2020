package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal - %v\n", e)
		os.Exit(1)
	}
}

func parseFile(s string) []string {
	// given a file, return a slice of strings,
	// each of which describes a row on the grid
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	a := make([]string, 0, 323)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	return a
}
