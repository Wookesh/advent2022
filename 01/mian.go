package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const test = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func main() {
	data, err := os.ReadFile("./advent2022/01/input.txt")
	if err != nil {
		log.Fatalf("os.ReadFile() failed: %v", err)
	}

	t := time.Now()
	resultOne := partOne(string(data))
	log.Printf("time: %v", time.Now().Sub(t))
	log.Printf("ans 1: %v", resultOne)

	t = time.Now()
	resultTwo := partTwo(string(data), 3)
	log.Printf("time: %v", time.Now().Sub(t))
	log.Printf("ans 2: %v", resultTwo)
}

func partOne(data string) int64 {
	most := int64(0)
	current := int64(0)
	for _, l := range strings.Split(data, "\n") {
		if l == "" {
			if current > most {
				most = current
			}
			current = 0
		} else {
			v, _ := strconv.ParseInt(l, 10, 64)
			current += v
		}
	}
	return most
}

func partTwo(data string, k int) int {
	var result []int
	current := 0
	for _, l := range strings.Split(data, "\n") {
		if l == "" {
			result = append(result, current)
			current = 0
		} else {
			v, _ := strconv.ParseInt(l, 10, 32)
			current += int(v)
		}
	}
	if current != 0 {
		result = append(result, current)
	}
	sort.Ints(result)
	total := 0
	for _, i := range result[len(result)-k:] {
		total += i
	}
	return total
}
