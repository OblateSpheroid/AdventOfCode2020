package main

import "fmt"

func leastWait(t int, ids []int) (int, int) {
	wait := t // start with high min wait time
	bus := -1
	for _, id := range ids {
		if id > 0 {
			tmp := (id - (t % id))
			if tmp < wait {
				wait = tmp
				bus = id
			}
		}
	}
	return bus, wait
}

func run(buses []int) int {
	for m := 1; true; m++ {
		good := true
		time := m * buses[0]
		for i, bus := range buses {
			if i == 0 || bus < 0 {
				continue
			}
			if (bus - (time % bus)) != i {
				good = false
				break
			}
		}
		if good {
			return time
		}
		if (time % 100) == 0 {
			fmt.Println(time)
		}
	}
	return -1
}

func main() {
	t, b := parseFile("test1.txt")
	bs, w := leastWait(t, b)
	fmt.Println((bs * w) == 295)
	fmt.Println(run(b) == 1068781)

	time, buses := parseFile("data.txt")
	bus, wait := leastWait(time, buses)
	fmt.Printf("Answer 1: %d\n", bus*wait)
	fmt.Printf("Answer 2: %d\n", run(buses))
}
