package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/impeccableai/impeccable/golibs/maps"
)

const test = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func main() {
	data, err := os.ReadFile("./advent2022/03/input.txt")
	if err != nil {
		log.Fatalf("os.ReadFile() failed: %v", err)
	}

	//data = []byte(test)

	t1 := time.Now()
	resultOne := partOne(string(data))
	log.Printf("time: %v", time.Now().Sub(t1))
	log.Printf("ans 1: %v", resultOne)

	t2 := time.Now()
	resultTwo := partTwo(string(data))
	log.Printf("time: %v", time.Now().Sub(t2))
	log.Printf("ans 2: %v", resultTwo)
}

func partOne(data string) int {
	prioritySum := 0
	for _, l := range strings.Split(data, "\n") {
		middle := len(l) / 2
		first, second := l[:middle], l[middle:]
		firstMapped := maps.NewFromSlice([]byte(first), func(r byte) byte { return r })
		for _, e := range []byte(second) {
			if _, ok := firstMapped[e]; !ok {
				continue
			}
			if e >= 'a' && e <= 'z' {
				prioritySum += int(e-'a') + 1
			} else {
				prioritySum += int(e-'A') + 27
			}
			break
		}
	}
	return prioritySum
}

func partTwo(data string) int {
	prioritySum := 0
	var group []string
	for _, l := range strings.Split(data, "\n") {
		group = append(group, l)
		if len(group) < 3 {
			continue
		}
		firstMapped := maps.NewFromSlice([]byte(group[0]), func(r byte) byte { return r })
		secondMapped := maps.NewFromSlice([]byte(group[1]), func(r byte) byte { return r })
		for _, c := range []byte(group[2]) {
			_, firstOk := firstMapped[c]
			_, secondOk := secondMapped[c]
			if !firstOk || !secondOk {
				continue
			}
			if c >= 'a' && c <= 'z' {
				prioritySum += int(c-'a') + 1
			} else {
				prioritySum += int(c-'A') + 27
			}
			break
		}
		group = nil
	}
	return prioritySum
}
