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

func parseFile(s string) (int, []int) {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	earliest, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	valid := []int{}
	for _, v := range buses {
		if v == "x" {
			valid = append(valid, -9)
		} else {
			i, _ := strconv.Atoi(v)
			valid = append(valid, i)
		}
	}

	return earliest, valid
}
