package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func main() {
	data, err := os.ReadFile("./advent2022/15/input.txt")
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

type Sensor struct {
	dist int
}

func partOne(s string) int {
	sensors := make(map[pos]*Sensor)
	beacons := make(map[pos]bool)
	for _, l := range strings.Split(s, "\n") {
		var x, y, ax, ay int
		_, err := fmt.Sscanf(l, "Sensor at x=%v, y=%v: closest beacon is at x=%v, y=%v", &x, &y, &ax, &ay)
		if err != nil {
			log.Fatal(err)
		}
		s := &Sensor{dist: abs(x-ax) + abs(y-ay)}
		sensors[pos{x, y}] = s
		beacons[pos{ax, ay}] = true
	}

	y := 2000000

	impossible := make(map[pos]bool)
	for p, s := range sensors {
		yDist := abs(p.y - y)
		distLeft := s.dist - yDist
		if distLeft <= 0 {
			continue
		}
		min := p.x - distLeft
		max := p.x + distLeft
		for i := min; i <= max; i++ {
			p := pos{i, y}
			if _, ok := beacons[p]; ok {
				continue
			}
			impossible[p] = true
		}
	}
	return len(impossible)
}

func partTwo(s string) int {
	sensors := make(map[pos]*Sensor)
	beacons := make(map[pos]bool)
	for _, l := range strings.Split(s, "\n") {
		var x, y, ax, ay int
		_, err := fmt.Sscanf(l, "Sensor at x=%v, y=%v: closest beacon is at x=%v, y=%v", &x, &y, &ax, &ay)
		if err != nil {
			log.Fatal(err)
		}
		s := &Sensor{dist: abs(x-ax) + abs(y-ay)}
		sensors[pos{x, y}] = s
		beacons[pos{ax, ay}] = true
	}

	toCheck := make(map[pos]bool)
	for p, s := range sensors {
		lDist := s.dist + 1
		for i := 0; i < s.dist; i++ {
			toCheck[pos{p.x + lDist - i, p.y + i}] = true
			toCheck[pos{p.x - lDist + i, p.y + i}] = true
			toCheck[pos{p.x + lDist - i, p.y - i}] = true
			toCheck[pos{p.x - lDist + i, p.y - i}] = true
		}
	}

	for p := range toCheck {
		if p.x < 0 || p.x > 4000000 || p.y < 0 || p.y > 4000000 {
			continue
		}
		bad := false
		for sp, s := range sensors {
			if abs(sp.x-p.x)+abs(sp.y-p.y) <= s.dist {
				bad = true
				break
			}
		}
		if !bad {
			return p.x*4000000 + p.y
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
