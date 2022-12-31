package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func main() {
	data, err := os.ReadFile("./advent2022/04/input.txt")
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

func partOne(s string) int {
	count := 0
	for _, l := range strings.Split(s, "\n") {
		var a, b, c, d int
		fmt.Sscanf(l, "%v-%v,%v-%v", &a, &b, &c, &d)
		if (a <= c && b >= d) || (c <= a && d >= b) {
			count++
		}
	}
	return count
}

func partTwo(s string) int {
	count := 0
	total := 0
	for _, l := range strings.Split(s, "\n") {
		var a, b, c, d int
		fmt.Sscanf(l, "%v-%v,%v-%v", &a, &b, &c, &d)
		if (a < c && b < c) || (d < a && d < b) {
			count++
		}
		total++
	}
	return total - count
}
