package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const (
	test = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
)

func main() {
	data, err := os.ReadFile("./advent2022/24/input.txt")
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

type pos struct {
	x, y int
}

func (p *pos) Move(m *move) pos {
	return pos{
		x: p.x + m.x,
		y: p.y + m.y,
	}
}

type posInTime struct {
	pos    pos
	minute int
}

type Blizard struct {
	direction move
}

var (
	directions = []*move{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{0, 0},
	}
)

func partOne(s string) int {
	blizzards := make(map[pos][]*Blizard)
	var end pos
	width := 0
	height := 0
	for x, s := range strings.Split(s, "\n") {
		width = len(s) - 1
		for y, c := range s {
			switch c {
			case '>':
				blizzards[pos{x, y}] = []*Blizard{{move{0, 1}}}
			case '<':
				blizzards[pos{x, y}] = []*Blizard{{move{0, -1}}}
			case '^':
				blizzards[pos{x, y}] = []*Blizard{{move{-1, 0}}}
			case 'v':
				blizzards[pos{x, y}] = []*Blizard{{move{1, 0}}}
			case '.':
				end = pos{x, y}
			}
		}
		height = x
	}
	start := pos{0, 1}

	log.Printf("start: %v, end: %v", start, end)
	log.Printf("width: %v, height: %v", width, height)

	bm := &BlizardsManager{
		blizardsInTime: map[int]map[pos][]*Blizard{0: blizzards},
		maxX:           height,
		maxY:           width,
	}

	queue := []posInTime{{start, 0}}
	seen := make(map[posInTime]bool)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		seen[posInTime{p.pos, p.minute}] = true
		//log.Printf("queue: %v, depth: %v", len(queue), len(bm.blizardsInTime))

		for _, dir := range directions {
			newP := p.pos.Move(dir)
			if newP == end {
				return p.minute + 1
			}

			if newP != start && (newP.x >= height || newP.x < 1 || newP.y < 1 || newP.y >= width) {
				continue
			}

			if _, ok := bm.getBlizzards(p.minute + 1)[newP]; ok {
				continue
			}

			if _, ok := seen[posInTime{newP, p.minute + 1}]; ok {
				continue
			}
			seen[posInTime{newP, p.minute + 1}] = true

			queue = append(queue, posInTime{newP, p.minute + 1})
		}
	}

	return 0
}

type BlizardsManager struct {
	blizardsInTime map[int]map[pos][]*Blizard
	maxX, maxY     int
}

func (b *BlizardsManager) getBlizzards(t int) map[pos][]*Blizard {
	currentBlizzard, ok := b.blizardsInTime[t]
	if !ok {
		currentBlizzard = map[pos][]*Blizard{}
		for pos, blizzards := range b.getBlizzards(t - 1) {
			for _, blizzard := range blizzards {
				if canProgress(&pos, blizzard, b.maxY, b.maxX) {
					newP := progress(&pos, blizzard)
					currentBlizzard[newP] = append(currentBlizzard[newP], blizzard)
				} else {
					newP := reset(&pos, blizzard, b.maxY, b.maxX)
					currentBlizzard[newP] = append(currentBlizzard[newP], blizzard)
				}
			}
		}
		//for i := 1; i < b.maxX; i++ {
		//	for j := 1; j < b.maxY; j++ {
		//		if b, ok := currentBlizzard[pos{i, j}]; ok {
		//			fmt.Print(len(b))
		//		} else {
		//			fmt.Print(".")
		//		}
		//	}
		//	fmt.Print("\n")
		//}
		//log.Printf("blizzard at: %v, count: %v", t, len(currentBlizzard))
		b.blizardsInTime[t] = currentBlizzard
	}
	return currentBlizzard
}

func canProgress(p *pos, b *Blizard, width, height int) bool {
	newP := p.Move(&b.direction)
	if newP.x >= height || newP.x < 1 || newP.y < 1 || newP.y >= width {
		return false
	}
	return true
}

func progress(p *pos, b *Blizard) pos {
	return p.Move(&b.direction)
}

func reset(p *pos, b *Blizard, width, height int) pos {
	if b.direction.x == -1 {
		return pos{height - 1, p.y}
	} else if b.direction.x == 1 {
		return pos{1, p.y}
	} else if b.direction.y == 1 {
		return pos{p.x, 1}
	} else {
		return pos{p.x, width - 1}
	}
}

func partTwo(s string) int {
	blizzards := make(map[pos][]*Blizard)
	var end pos
	width := 0
	height := 0
	for x, s := range strings.Split(s, "\n") {
		width = len(s) - 1
		for y, c := range s {
			switch c {
			case '>':
				blizzards[pos{x, y}] = []*Blizard{{move{0, 1}}}
			case '<':
				blizzards[pos{x, y}] = []*Blizard{{move{0, -1}}}
			case '^':
				blizzards[pos{x, y}] = []*Blizard{{move{-1, 0}}}
			case 'v':
				blizzards[pos{x, y}] = []*Blizard{{move{1, 0}}}
			case '.':
				end = pos{x, y}
			}
		}
		height = x
	}
	start := pos{0, 1}

	bm := &BlizardsManager{
		blizardsInTime: map[int]map[pos][]*Blizard{0: blizzards},
		maxX:           height,
		maxY:           width,
	}

	m := calculateMoves(bm, width, height, start, end, 0)
	log.Printf("first part: %v", m)
	m = calculateMoves(bm, width, height, end, start, m)
	log.Printf("return part: %v", m)
	m = calculateMoves(bm, width, height, start, end, m)
	log.Printf("goal part: %v", m)

	return m
}

func calculateMoves(bm *BlizardsManager, width, height int, start, end pos, t int) int {
	queue := []posInTime{{start, t}}
	seen := make(map[posInTime]bool)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		seen[posInTime{p.pos, p.minute}] = true
		//log.Printf("queue: %v, depth: %v", len(queue), len(bm.blizardsInTime))

		for _, dir := range directions {
			newP := p.pos.Move(dir)
			if newP == end {
				return p.minute + 1
			}

			if newP != start && (newP.x >= height || newP.x < 1 || newP.y < 1 || newP.y >= width) {
				continue
			}

			if _, ok := bm.getBlizzards(p.minute + 1)[newP]; ok {
				continue
			}

			if _, ok := seen[posInTime{newP, p.minute + 1}]; ok {
				continue
			}
			seen[posInTime{newP, p.minute + 1}] = true

			queue = append(queue, posInTime{newP, p.minute + 1})
		}
	}
	return 0
}
