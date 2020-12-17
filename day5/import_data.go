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
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	a := make([]string, 0, 868)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	return a
}
