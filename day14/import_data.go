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

func parseFile(s string) []Card {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	cards := []Card{}
	i := -1
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		sl := strings.Split(line, " = ")
		if sl[0][1] == 'a' { // mask instruction
			i++
			cards = append(cards, Card{})
			cards[i].mask = sl[1]
		}
		if sl[0][1] == 'e' { // mem instruction
			tmp := strings.ReplaceAll(sl[0], "mem[", "")
			mem, _ := strconv.Atoi(strings.ReplaceAll(tmp, "]", ""))
			val, _ := strconv.Atoi(sl[1])
			cards[i].inst = append(cards[i].inst, [2]int{mem, val})
		}
	}

	return cards
}
