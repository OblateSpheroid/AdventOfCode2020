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

func parseFile(s string) [][]string {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	cs := [][]string{}
	c := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			cs = append(cs, c) // finished with current card, append it
			c = []string{}     // reset c
			continue           // don't need to do anything else
		}
		c = append(c, line) // append line to the current card
	}

	return cs
}
