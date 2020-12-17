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

func parseFile(s string) Rules {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	rules := make(Rules)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		outter := strings.Split(line, " bags contain ")
		subj := outter[0]
		rules[subj] = make(map[string]int)
		objs := strings.Split(outter[1], ", ")
		for _, v := range objs {
			tmp := strings.Split(v, " ")
			obj := tmp[1] + " " + tmp[2]
			if tmp[0] != "no" {
				rules[subj][obj], _ = strconv.Atoi(tmp[0])
			} else {
				rules[subj]["other"] = 0
			}
		}
	}
	return rules
}
