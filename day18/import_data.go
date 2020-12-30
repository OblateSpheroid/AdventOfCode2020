package main

import (
	"aoc2020/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal - %v\n", e)
		os.Exit(1)
	}
}

func parseFile(s string) [][]string {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	a := [][]string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		line = helpers.DropString(line, " ") // remove empty spaces
		a = append(a, line)
	}

	return a
}
