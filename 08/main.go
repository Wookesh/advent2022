package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const test = `30373
25512
65332
33549
35390`

func main() {
	data, err := os.ReadFile("./advent2022/08/input.txt")
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

type move struct {
	x, y int
}

var (
	up    = move{0, 1}
	down  = move{0, -1}
	right = move{1, 0}
	left  = move{-1, 0}

	directions = []move{up, down, right, left}
)

func partOne(s string) int {
	var trees [][]int
	for _, l := range strings.Split(s, "\n") {
		var treeLine []int
		for _, c := range l {
			treeLine = append(treeLine, int(c-'0'))
		}
		trees = append(trees, treeLine)
	}

	totalVisible := 0
	for x, l := range trees {
		for y, _ := range l {
			visible := false
			for _, dir := range directions {
				if allSmaller(trees, x, y, dir) {
					visible = true
					break
				}
			}
			if visible {
				totalVisible += 1
			}
		}
	}

	return totalVisible
}

func allSmaller(trees [][]int, x, y int, dir move) bool {
	for i, j := x+dir.x, y+dir.y; i < len(trees) && i >= 0 && j < len(trees[i]) && j >= 0; i, j = i+dir.x, j+dir.y {
		if trees[i][j] >= trees[x][y] {
			return false
		}
	}
	return true
}

func partTwo(s string) int {
	var trees [][]int
	for _, l := range strings.Split(s, "\n") {
		var treeLine []int
		for _, c := range l {
			treeLine = append(treeLine, int(c-'0'))
		}
		trees = append(trees, treeLine)
	}

	bestScore := 0
	for x, l := range trees {
		for y, _ := range l {
			score := 1
			for _, dir := range directions {
				s := visibleTrees(trees, x, y, dir)
				score = score * s
			}
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func visibleTrees(trees [][]int, x, y int, dir move) int {
	visible := 0
	for i, j := x+dir.x, y+dir.y; i < len(trees) && i >= 0 && j < len(trees[i]) && j >= 0; i, j = i+dir.x, j+dir.y {
		visible++
		if trees[i][j] >= trees[x][y] {
			break
		}
	}
	return visible
}
