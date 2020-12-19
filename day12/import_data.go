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

func parseFile(s string) []Instruction {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	a := []Instruction{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line[1:])
		checkError(err)
		inst := Instruction{order: rune(line[0]), num: n}
		a = append(a, inst)
	}

	return a
}
