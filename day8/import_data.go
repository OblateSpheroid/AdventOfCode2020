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

func parseFile(s string) []Cmd {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	cmds := []Cmd{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		cmd := Cmd{}
		line := strings.Split(scanner.Text(), " ")
		cmd.op = line[0]
		cmd.arg, _ = strconv.Atoi(line[1])
		cmds = append(cmds, cmd)
	}

	return cmds
}
