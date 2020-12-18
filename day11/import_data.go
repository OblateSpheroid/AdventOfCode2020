package main

/* Read in data from file and export as a grid of seats */

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

func parseFile(s string) Seats {
	f, fileErr := os.Open(s)
	defer f.Close()
	checkError(fileErr)

	seats := Seats{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Seat, len(line))
		for i, v := range line {
			if v == 'L' {
				row[i] = Seat{seat: true}
			} else if v == '.' {
				row[i] = Seat{seat: false}
			} else {
				checkError(fmt.Errorf("%c is neither L nor .", v))
			}
		}
		seats.appendInPlace(row)
	}
	return seats
}
