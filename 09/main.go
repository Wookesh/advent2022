package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func main() {
	data, err := os.ReadFile("./advent2022/09/input.txt")
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

type position struct {
	x, y int
}

func (p *position) move(m move) position {
	return position{p.x + m.x, p.y + m.y}
}

type move struct {
	x, y int
}

func getMove(h, t position) move {
	if h.x == t.x {
		return move{0, (h.y - t.y) / abs(h.y-t.y)}
	}
	if h.y == t.y {
		return move{(h.x - t.x) / abs(h.x-t.x), 0}
	}
	return move{(h.x - t.x) / abs(h.x-t.x), (h.y - t.y) / abs(h.y-t.y)}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dist(a, b position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func touching(a, b position) bool {
	return abs(a.x-b.x) <= 1 && abs(a.y-b.y) <= 1
}

func sameRowOrColumn(h, t position) bool {
	return h.x == t.x || h.y == t.y
}

func partOne(s string) int {
	var tail, head position

	visited := make(map[position]bool)
	visited[tail] = true

	for _, l := range strings.Split(s, "\n") {
		var direction string
		var count int
		fmt.Sscanf(l, "%s %v", &direction, &count)

		for i := 0; i < count; i++ {

			switch direction {
			case "R":
				head.x += 1
			case "U":
				head.y += 1
			case "L":
				head.x -= 1
			case "D":
				head.y -= 1
			}

			if sameRowOrColumn(head, tail) && dist(head, tail) >= 2 {
				tail = tail.move(getMove(head, tail))
			} else if !touching(head, tail) {
				tail = tail.move(getMove(head, tail))
			}

			visited[tail] = true
		}
	}

	return len(visited)
}

func partTwo(s string) int {
	knots := make([]position, 10)
	lastKnot := len(knots) - 1

	visited := make(map[position]bool)
	visited[knots[lastKnot]] = true

	for _, l := range strings.Split(s, "\n") {
		var direction string
		var count int
		fmt.Sscanf(l, "%s %v", &direction, &count)

		for i := 0; i < count; i++ {

			switch direction {
			case "R":
				knots[0].x += 1
			case "U":
				knots[0].y += 1
			case "L":
				knots[0].x -= 1
			case "D":
				knots[0].y -= 1
			}

			for i := 0; i < len(knots)-1; i++ {
				if sameRowOrColumn(knots[i], knots[i+1]) && dist(knots[i], knots[i+1]) >= 2 {
					knots[i+1] = knots[i+1].move(getMove(knots[i], knots[i+1]))
				} else if !touching(knots[i], knots[i+1]) {
					knots[i+1] = knots[i+1].move(getMove(knots[i], knots[i+1]))
				}
			}

			visited[knots[lastKnot]] = true
		}
	}

	return len(visited)
}
