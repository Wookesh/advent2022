package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	test = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

	test0 = `1,1,1
2,1,1`
)

func main() {
	data, err := os.ReadFile("./advent2022/18/input.txt")
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

type cube struct {
	p pos
}

type pos struct {
	x, y, z int
}

func (p *pos) Move(m move) pos {
	return pos{
		x: p.x + m.x,
		y: p.y + m.y,
		z: p.z + m.z,
	}
}

type move struct {
	x, y, z int
}

var (
	directions = []move{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}
)

func partOne(s string) int {
	cubes := make(map[pos]*cube)
	for _, l := range strings.Split(s, "\n") {
		coords := strings.Split(l, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		p := pos{
			x: x,
			y: y,
			z: z,
		}
		cubes[p] = &cube{p}
	}
	total := 0
	for _, c := range cubes {
		//log.Printf("%#v", c)
		for _, d := range directions {
			np := c.p.Move(d)
			if _, ok := cubes[np]; ok {
				continue
			}
			total += 1
		}
	}

	return total
}

func partTwo(s string) int {
	cubes := make(map[pos]*cube)
	for _, l := range strings.Split(s, "\n") {
		coords := strings.Split(l, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		p := pos{
			x: x,
			y: y,
			z: z,
		}
		cubes[p] = &cube{p}
	}

	maxX := 0
	maxY := 0
	maxZ := 0
	for _, c := range cubes {
		if c.p.x > maxX {
			maxX = c.p.x
		}
		if c.p.y > maxY {
			maxY = c.p.y
		}
		if c.p.z > maxZ {
			maxZ = c.p.z
		}
	}

	total := 0
	rm := make(map[pos]bool)
	for _, c := range cubes {
		for _, d := range directions {
			np := c.p.Move(d)
			rv, ok := rm[np]
			if !ok {
				rv = reachable(np, cubes, pos{maxX + 1, maxY + 1, maxZ + 1})
				rm[np] = rv
			}
			if !rv {
				continue
			}

			if _, ok := cubes[np]; ok {
				continue
			}
			total += 1
		}
	}

	return total
}

func reachable(dst pos, cubes map[pos]*cube, start pos) bool {
	queue := []pos{start}
	seen := make(map[pos]bool)

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		seen[q] = true

		for _, d := range directions {
			p := q.Move(d)
			if p == dst {
				return true
			}
			if _, ok := seen[p]; ok {
				continue
			}
			if p.x > start.x || p.x < -1 {
				continue
			}
			if p.y > start.y || p.y < -1 {
				continue
			}
			if p.z > start.z || p.z < -1 {
				continue
			}
			if _, ok := cubes[p]; ok {
				continue
			}

			seen[p] = true
			queue = append(queue, p)
		}
	}
	return false
}
