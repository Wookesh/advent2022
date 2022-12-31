package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const test = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func main() {
	data, err := os.ReadFile("./advent2022/14/input.txt")
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

type pos struct {
	x, y int
}

type elem int8

const (
	empty elem = iota
	block
	sand
)

func partOne(s string) int {
	var moreLines [][]pos
	maxX, maxY := 0, 0
	firstX, firstY := 500, 500
	for _, l := range strings.Split(s, "\n") {
		parts := strings.Split(l, " -> ")
		var lines []pos
		for _, p := range parts {
			xs, ys, _ := strings.Cut(p, ",")
			x, _ := strconv.Atoi(xs)
			y, _ := strconv.Atoi(ys)
			x, y = y, x
			lines = append(lines, pos{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			if x < firstX {
				firstX = x
			}
			if y < firstY {
				firstY = y
			}
		}
		moreLines = append(moreLines, lines)
	}

	var board [][]elem
	for i := 0; i <= maxX; i++ {
		var line []elem
		for j := 0; j <= maxY; j++ {
			line = append(line, empty)
		}
		board = append(board, line)
	}

	for _, line := range moreLines {
		for i := 1; i < len(line); i++ {
			if line[i-1].x == line[i].x {
				a, b := line[i-1].y, line[i].y
				if a > b {
					a, b = b, a
				}
				for y := a; y <= b; y++ {
					board[line[i-1].x][y] = block
				}
			}
			if line[i-1].y == line[i].y {
				a, b := line[i-1].x, line[i].x
				if a > b {
					a, b = b, a
				}
				for x := a; x <= b; x++ {
					board[x][line[i].y] = block
				}
			}
		}
	}

	source := pos{0, 500}

	for i := 0; ; i++ {
		fellOut := addSand(board, source)
		if fellOut {
			return i
		}
	}
}

func addSand(board [][]elem, source pos) bool {
	s := source
	for {
		down := pos{s.x + 1, s.y}
		if down.x >= len(board) {
			return true
		}
		if board[down.x][down.y] == empty {
			s = down
			continue
		}
		if board[down.x][down.y-1] == empty {
			s = pos{down.x, down.y - 1}
			continue
		}
		if board[down.x][down.y+1] == empty {
			s = pos{down.x, down.y + 1}
			continue
		}
		board[s.x][s.y] = sand
		return false
	}
}

func partTwo(s string) int {
	var moreLines [][]pos
	maxX, maxY := 0, 0
	firstX, firstY := 500, 500
	for _, l := range strings.Split(s, "\n") {
		parts := strings.Split(l, " -> ")
		var lines []pos
		for _, p := range parts {
			xs, ys, _ := strings.Cut(p, ",")
			x, _ := strconv.Atoi(xs)
			y, _ := strconv.Atoi(ys)
			x, y = y, x
			lines = append(lines, pos{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			if x < firstX {
				firstX = x
			}
			if y < firstY {
				firstY = y
			}
		}
		moreLines = append(moreLines, lines)
	}

	moreLines = append(moreLines, []pos{{maxX + 2, 0}, {maxX + 2, 1000}})
	maxY = 1000
	maxX = maxX + 2

	var board [][]elem
	for i := 0; i <= maxX; i++ {
		var line []elem
		for j := 0; j <= maxY; j++ {
			line = append(line, empty)
		}
		board = append(board, line)
	}

	for _, line := range moreLines {
		for i := 1; i < len(line); i++ {
			if line[i-1].x == line[i].x {
				a, b := line[i-1].y, line[i].y
				if a > b {
					a, b = b, a
				}
				for y := a; y <= b; y++ {
					board[line[i-1].x][y] = block
				}
			}
			if line[i-1].y == line[i].y {
				a, b := line[i-1].x, line[i].x
				if a > b {
					a, b = b, a
				}
				for x := a; x <= b; x++ {
					board[x][line[i].y] = block
				}
			}
		}
	}

	source := pos{0, 500}

	for i := 0; ; i++ {
		fellOut := addSand2(board, source)
		if fellOut {
			return i + 1
		}
	}
}

func addSand2(board [][]elem, source pos) bool {
	s := source
	for {
		down := pos{s.x + 1, s.y}
		if down.x >= len(board) {
			log.Fatalf("out of board: %v", down)
			return true
		}
		if board[down.x][down.y] == empty {
			s = down
			continue
		}
		if board[down.x][down.y-1] == empty {
			s = pos{down.x, down.y - 1}
			continue
		}
		if board[down.x][down.y+1] == empty {
			s = pos{down.x, down.y + 1}
			continue
		}
		board[s.x][s.y] = sand
		return board[source.x][source.y] == sand
	}
}
