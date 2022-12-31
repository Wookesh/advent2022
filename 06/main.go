package main

import (
	"log"
	"os"
	"time"
)

const test = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func main() {
	data, err := os.ReadFile("./advent2022/06/input.txt")
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
	for i := range s {
		if i < 4 {
			continue
		}
		m := make(map[rune]bool)
		for _, k := range s[i-4 : i] {
			m[k] = true
		}
		if len(m) == 4 {
			return i
		}
	}
	return 0
}

func partTwo(s string) int {
	for i := range s {
		if i < 14 {
			continue
		}
		m := make(map[rune]bool)
		for _, k := range s[i-14 : i] {
			m[k] = true
		}
		if len(m) == 14 {
			return i
		}
	}
	return 0
}
