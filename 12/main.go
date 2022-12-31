package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const test = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func main() {
	data, err := os.ReadFile("./advent2022/12/input.txt")
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

func (p *pos) move(m *move) pos {
	return pos{p.x + m.x, p.y + m.y}
}

type move struct {
	x, y int
}

var (
	directions = []*move{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
)

type posWithDist struct {
	p    pos
	dist int
}

func partOne(s string) int {
	var m [][]int
	var start, end pos
	maxY := 0
	for i, l := range strings.Split(s, "\n") {
		var line []int
		for j, c := range l {
			if c == 'S' {
				start = pos{i, j}
				log.Printf("start: %v", start)
				line = append(line, 0)
			} else if c == 'E' {
				end = pos{i, j}
				log.Printf("end: %v", end)
				line = append(line, int('z'-'a'))
			} else {
				line = append(line, int(c-'a'))
			}
		}
		maxY = len(line)
		m = append(m, line)
	}

	visited := make(map[pos]bool)
	var queue []posWithDist
	queue = append(queue, posWithDist{start, 0})
	visited[start] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		hV := m[v.p.x][v.p.y]

		for _, dir := range directions {
			newPos := v.p.move(dir)
			if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(m) || newPos.y >= maxY {
				continue
			}
			hn := m[newPos.x][newPos.y]
			if hn > hV+1 {
				continue
			}
			if newPos == end {
				return v.dist + 1
			}
			if _, ok := visited[newPos]; ok {
				continue
			}
			visited[newPos] = true
			queue = append(queue, posWithDist{newPos, v.dist + 1})
		}
	}

	return 0
}

func partTwo(s string) int {
	var m [][]int
	var start, end pos
	var possibleStarting []pos
	maxY := 0
	for i, l := range strings.Split(s, "\n") {
		var line []int
		for j, c := range l {
			if c == 'S' {
				start = pos{i, j}
				log.Printf("start: %v", start)
				possibleStarting = append(possibleStarting, start)
				line = append(line, 0)
			} else if c == 'E' {
				end = pos{i, j}
				log.Printf("end: %v", end)
				line = append(line, int('z'-'a'))
			} else {
				if c == 'a' {
					possibleStarting = append(possibleStarting, pos{i, j})
				}
				line = append(line, int(c-'a'))
			}
		}
		maxY = len(line)
		m = append(m, line)
	}

	shortest := 500
	for _, s := range possibleStarting {
		result, found := shortestPath(m, s, end, maxY)
		if !found {
			continue
		}
		if result < shortest {
			shortest = result
		}
	}

	return shortest
}

func shortestPath(m [][]int, start, end pos, maxY int) (int, bool) {
	visited := make(map[pos]bool)
	var queue []posWithDist
	queue = append(queue, posWithDist{start, 0})
	visited[start] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		hV := m[v.p.x][v.p.y]

		for _, dir := range directions {
			newPos := v.p.move(dir)
			if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(m) || newPos.y >= maxY {
				continue
			}
			hn := m[newPos.x][newPos.y]
			if hn > hV+1 {
				continue
			}

			if newPos == end {
				return v.dist + 1, true
			}
			if _, ok := visited[newPos]; ok {
				continue
			}
			visited[newPos] = true

			queue = append(queue, posWithDist{newPos, v.dist + 1})
		}
	}
	return 0, false
}
