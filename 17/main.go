package main

import (
	"log"
	"os"
	"time"
)

const test = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

var (
	blocks = [][][]byte{
		[][]byte{
			[]byte("####"),
		},
		[][]byte{
			[]byte(".#."),
			[]byte("###"),
			[]byte(".#."),
		},
		[][]byte{
			[]byte("..#"),
			[]byte("..#"),
			[]byte("###"),
		},
		[][]byte{
			[]byte("#"),
			[]byte("#"),
			[]byte("#"),
			[]byte("#"),
		},
		[][]byte{
			[]byte("##"),
			[]byte("##"),
		},
	}
)

func main() {
	data, err := os.ReadFile("./advent2022/17/input.txt")
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

type Board struct {
	b [][]byte

	currentBlock      [][]byte
	currentBlockIndex int
	currentBlockTop   int
	currentBlockX     int
}

func (b *Board) simulate(c rune) bool {
	if b.currentBlock == nil {
		b.currentBlock = blocks[b.currentBlockIndex]
		currentTopA := findTop(b.b)
		_ = currentTopA
		b.extend()
		currentTop := findTop(b.b)
		b.currentBlockTop = currentTop + len(b.currentBlock) + 3 - 1
		b.currentBlockX = 2
	}

	if b.canMove(c) {
		b.move(c)
	}

	if b.canFall() {
		b.fall()
		return false
	} else {
		b.save()
		b.currentBlock = nil
		b.currentBlockIndex = (b.currentBlockIndex + 1) % len(blocks)
		return true
	}
}

func (b *Board) save() {
	for i, row := range b.currentBlock {
		for j, k := range row {
			if k == '.' {
				continue
			}
			b.b[b.currentBlockTop-i][b.currentBlockX+j] = '#'
		}
	}
}

func (b *Board) canMove(c rune) bool {
	switch c {
	case '<':
		if b.currentBlockX-1 < 0 {
			return false
		}
		for i, row := range b.currentBlock {
			for j, k := range row {
				if k == '#' && b.b[b.currentBlockTop-i][b.currentBlockX+j-1] == '#' {
					return false
				}
			}
		}
	case '>':
		if b.currentBlockX+1+len(b.currentBlock[0]) > len(b.b[0]) {
			return false
		}
		for i, row := range b.currentBlock {
			for j, k := range row {
				if k == '#' && b.b[b.currentBlockTop-i][b.currentBlockX+j+1] == '#' {
					return false
				}
			}
		}
	}
	return true
}
func (b *Board) move(c rune) {
	switch c {
	case '<':
		b.currentBlockX -= 1
	case '>':
		b.currentBlockX += 1
	}
}

func (b *Board) canFall() bool {
	if b.currentBlockTop-len(b.currentBlock) < 0 {
		return false
	}
	for i, row := range b.currentBlock {
		for j, k := range row {
			if b.b[b.currentBlockTop-i-1][b.currentBlockX+j] == '#' && k == '#' {
				return false
			}
		}
	}
	return true
}

func (b *Board) fall() {
	b.currentBlockTop -= 1
}

func (b *Board) extend() {
	top := findTop(b.b)
	for len(b.b)-top <= len(b.currentBlock)+3 {
		b.b = append(b.b, newRow())
	}
}

func (b *Board) String() string {
	s := ""
	for i := 0; i < len(b.b); i++ {
		y := len(b.b) - 1 - i
		for j, c := range b.b[y] {
			blockTop := b.currentBlockTop
			blockBottom := b.currentBlockTop - len(b.currentBlock) + 1
			blockStart := b.currentBlockX
			blockEnd := b.currentBlockX + len(b.currentBlock[0]) - 1
			blockY := b.currentBlockTop - y
			blockX := j - b.currentBlockX
			if b.currentBlock != nil &&
				blockBottom <= y && y <= blockTop &&
				blockStart <= j && j <= blockEnd &&
				b.currentBlock[blockY][blockX] == '#' {
				s += string('@')
			} else {
				s += string(c)
			}
		}
		s += "\n"
	}
	s += "\n"
	return s
}

func (b *Board) describeTop() [7]int {
	top := findTop(b.b)
	tops := make(map[int]int)
	for i := top; i > 0; i-- {
		for i, c := range b.b[top] {
			if _, ok := tops[i]; ok {
				continue
			}
			if c == '#' {
				tops[i] = top - i
			}
		}
		if len(tops) == 7 {
			result := [7]int{}
			for i, c := range tops {
				result[i] = c
			}
			return result
		}
	}
	return [7]int{}
}

func newRow() []byte {
	return []byte(".......")
}

func findTop(board [][]byte) int {
	for i := 0; i < len(board); i++ {
		for _, c := range board[len(board)-1-i] {
			if c == '#' {
				return len(board) - i
			}
		}
	}
	return 0
}

func partOne(s string) int {
	board := &Board{}

	blocksFall := 0
	for ci := 0; ; ci = (ci + 1) % len(s) {
		fall := board.simulate(rune([]byte(s)[ci]))
		if fall {
			blocksFall += 1
		}
		if blocksFall == 2022 {
			break
		}
	}

	return findTop(board.b)
}

type state struct {
	top [7]int
	ci  int
	bi  int
}

type stateVal struct {
	i      int
	blocks int
	top    int
}

func partTwo(s string) int {
	board := &Board{}

	total := 1000000000000
	left := total

	blocksFall := 0
	cycleFirstTop := 0
	blocksCountCycle := 0
	extraTop := 0
	seen := make(map[state]stateVal)
	for i, ci := 0, 0; ; i, ci = i+1, (ci+1)%len(s) {
		fall := board.simulate(rune([]byte(s)[ci]))
		if fall {
			blocksFall += 1
			if extraTop == 0 {
				state := state{
					top: board.describeTop(),
					ci:  ci,
					bi:  board.currentBlockIndex,
				}
				if val, ok := seen[state]; ok {
					blocksCountCycle = blocksFall - val.blocks
					if (total-blocksFall)%blocksCountCycle != 0 {
						continue
					}
					cycleFirstTop = val.top
					log.Printf("found cycle: with %v blocks, after %v blocks, firstOccurrenceTop: %v", blocksCountCycle, blocksFall, cycleFirstTop)
					left = total - blocksFall
					log.Printf("currentlyLeft: %v", left)
					times := left / blocksCountCycle
					log.Printf("times: %v", times)
					topChangePerCycle := findTop(board.b) - cycleFirstTop
					log.Printf("topChangePerCycle: %v", topChangePerCycle)
					extraTop = times * topChangePerCycle
					log.Printf("extra top: %v", extraTop)
					blocksFall += times * blocksCountCycle
					log.Printf("blocksFall: %v", blocksFall)
					log.Printf("total - blocksFall: %v", total-blocksFall)
				}
				seen[state] = stateVal{i: i, top: findTop(board.b), blocks: blocksFall}
			}
		}
		if blocksFall == total {
			break
		}
	}

	// 1514285714288
	return findTop(board.b) + extraTop
}

func partTwoDumb(s string) int {
	board := &Board{}

	blocksFall := 0
	for ci := 0; ; ci = (ci + 1) % len(s) {
		fall := board.simulate(rune([]byte(s)[ci]))
		if fall {
			blocksFall += 1
		}
		if blocksFall == 1000000000000 {
			break
		}
	}

	return findTop(board.b)
}
