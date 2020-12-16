package main

import (
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

func parseFile(s string) []Passport {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	var ps []Passport
	p := make(Passport)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ps = append(ps, p) // finished with current passport, append it
			p = make(Passport) // reset p
			continue           // don't need to do anything else
		}
		tmp := strings.Split(line, " ")
		for _, v := range tmp {
			pair := strings.SplitN(v, ":", 2)
			p[pair[0]] = pair[1]
		}
	}

	return ps
}
