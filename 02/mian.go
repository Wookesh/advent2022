package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const test = `A Y
B X
C Z`

func main() {
	data, err := os.ReadFile("./advent2022/02/input.txt")
	if err != nil {
		log.Fatalf("os.ReadFile() failed: %v", err)
	}

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
	score := 0
	for _, l := range strings.Split(data, "\n") {
		a, b := string([]byte(l)[0]), string([]byte(l)[2])
		switch b {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}
		if a == "A" && b == "X" ||
			a == "B" && b == "Y" ||
			a == "C" && b == "Z" {
			score += 3
		}
		if a == "A" && b == "Y" ||
			a == "B" && b == "Z" ||
			a == "C" && b == "X" {
			score += 6
		}

	}
	return score
}

func partTwo(data string) int {
	score := 0
	for _, l := range strings.Split(data, "\n") {
		a, b := string([]byte(l)[0]), string([]byte(l)[2])
		switch b {
		case "X":
			score += 0
			switch a {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		case "Y":
			score += 3
			switch a {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		case "Z":
			score += 6
			switch a {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}
	}
	return score
}
