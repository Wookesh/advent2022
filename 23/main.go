package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	test = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

	test2 = `.....
..##.
..#..
.....
..##.
.....`
)

func main() {
	data, err := os.ReadFile("./advent2022/23/input.txt")
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

type pos struct {
	x, y int
}

func partOne(s string) int {
	elves := make(map[pos]bool)
	for i, l := range strings.Split(s, "\n") {
		for j, c := range l {
			if c == '#' {
				elves[pos{x: i, y: j}] = true
			}
		}
	}

	var moveProps []func(map[pos]pos, pos, near) bool
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.N && !n.NW && !n.NE {
			propositions[e] = pos{e.x - 1, e.y}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.S && !n.SW && !n.SE {
			propositions[e] = pos{e.x + 1, e.y}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.W && !n.NW && !n.SW {
			propositions[e] = pos{e.x, e.y - 1}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.E && !n.NE && !n.SE {
			propositions[e] = pos{e.x, e.y + 1}
			return true
		}
		return false
	})

	for i := 0; i < 10; i++ {
		elves, _ = simulate(elves, moveProps)
		firstMove := moveProps[0]
		moveProps = append(moveProps[1:], firstMove)
	}

	return countEmpty(elves)
}

type near struct {
	N, NW, NE, S, SW, SE, W, E bool
}

func simulate(elves map[pos]bool, moveProps []func(map[pos]pos, pos, near) bool) (map[pos]bool, bool) {
	propositions := make(map[pos]pos)
	for e := range elves {
		var n near
		_, n.N = elves[pos{e.x - 1, e.y}]
		_, n.NW = elves[pos{e.x - 1, e.y - 1}]
		_, n.NE = elves[pos{e.x - 1, e.y + 1}]
		_, n.S = elves[pos{e.x + 1, e.y}]
		_, n.SW = elves[pos{e.x + 1, e.y - 1}]
		_, n.SE = elves[pos{e.x + 1, e.y + 1}]
		_, n.W = elves[pos{e.x, e.y - 1}]
		_, n.E = elves[pos{e.x, e.y + 1}]

		if !n.N && !n.NW && !n.NE && !n.S && !n.SW && !n.SE && !n.W && !n.E {
			continue
		}

		for _, f := range moveProps {
			if f(propositions, e, n) {
				break
			}
		}
	}

	sameProps := make(map[pos]int)
	for _, prop := range propositions {
		sameProps[prop] += 1
	}

	anyMoved := false

	newElves := make(map[pos]bool)
	for e := range elves {
		if prop, ok := propositions[e]; !ok || prop == e {
			newElves[e] = true
			continue
		}
		if sameProps[propositions[e]] > 1 {
			newElves[e] = true
			continue
		}
		anyMoved = true
		newElves[propositions[e]] = true
	}

	return newElves, anyMoved
}

func countEmpty(elves map[pos]bool) int {
	minX, maxX, minY, maxY := getBoundraies(elves)

	//log.Printf("elves: %v, x: %v : %v, y: %v : %v", len(elves), minX, maxX, minY, maxY)
	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func getBoundraies(elves map[pos]bool) (minX, maxX, minY, maxY int) {
	init := false
	for p := range elves {
		if !init {
			minX, maxX, minY, maxY = p.x, p.x, p.y, p.y
			init = true
		}
		if minX > p.x {
			minX = p.x
		}
		if maxX < p.x {
			maxX = p.x
		}
		if minY > p.y {
			minY = p.y
		}
		if maxY < p.y {
			maxY = p.y
		}
	}
	return
}

func printElves(elves map[pos]bool) {
	minX, maxX, minY, maxY := getBoundraies(elves)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if _, ok := elves[pos{i, j}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func partTwo(s string) int {
	elves := make(map[pos]bool)
	for i, l := range strings.Split(s, "\n") {
		for j, c := range l {
			if c == '#' {
				elves[pos{x: i, y: j}] = true
			}
		}
	}

	var moveProps []func(map[pos]pos, pos, near) bool
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.N && !n.NW && !n.NE {
			propositions[e] = pos{e.x - 1, e.y}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.S && !n.SW && !n.SE {
			propositions[e] = pos{e.x + 1, e.y}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.W && !n.NW && !n.SW {
			propositions[e] = pos{e.x, e.y - 1}
			return true
		}
		return false
	})
	moveProps = append(moveProps, func(propositions map[pos]pos, e pos, n near) bool {
		if !n.E && !n.NE && !n.SE {
			propositions[e] = pos{e.x, e.y + 1}
			return true
		}
		return false
	})

	i := 1
	for ; ; i++ {
		anyMoved := false
		elves, anyMoved = simulate(elves, moveProps)
		if !anyMoved {
			break
		}
		firstMove := moveProps[0]
		moveProps = append(moveProps[1:], firstMove)
	}

	return i
}
