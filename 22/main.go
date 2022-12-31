package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	test = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
)

func main() {
	data, err := os.ReadFile("./advent2022/22/input.txt")
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

func partOne(s string) int {
	m, moves, ok := strings.Cut(s, "\n\n")
	if !ok {
		panic("not ok")
	}
	var board [][]rune
	maxWidth := 0
	for _, l := range strings.Split(m, "\n") {
		var boardLine []rune
		for _, c := range l {
			boardLine = append(boardLine, c)
		}
		if len(boardLine) > maxWidth {
			maxWidth = len(boardLine)
		}
		board = append(board, boardLine)
	}

	for i := range board {
		for len(board[i]) < maxWidth {
			board[i] = append(board[i], ' ')
		}
	}

	p := pos{0, findStartY(board), 0}
	for _, m := range parseMoves(moves) {
		if m.rotation != 0 {
			p.direction += m.rotation
			if p.direction < 0 {
				p.direction = len(directions) - 1
			}
			p.direction = p.direction % len(directions)
		} else {
			for i := 0; i < m.forward; i++ {
				moved := false
				p, moved = p.move(board)
				if !moved {
					break
				}
			}
		}
	}

	return 1000*(p.x+1) + 4*(p.y+1) + p.direction
}

func at(b [][]rune, p pos) rune {
	return b[p.x][p.y]
}

func findStartY(board [][]rune) int {
	for i, c := range board[0] {
		if c == '.' {
			return i
		}
	}
	return 0
}

var (
	right  = direction{0, 1}
	down   = direction{1, 0}
	left   = direction{0, -1}
	up     = direction{-1, 0}
	rightI = 0
	downI  = 1
	leftI  = 2
	upI    = 3

	directions = []direction{right, down, left, up}
)

type pos struct {
	x, y      int
	direction int
}

func (p *pos) move(board [][]rune) (pos, bool) {
	nx := p.x + directions[p.direction].x
	ny := p.y + directions[p.direction].y
	if directions[p.direction].x == 1 {
		if nx >= len(board) || board[nx][ny] == ' ' {
			for i := 0; i < len(board); i++ {
				if board[i][p.y] == ' ' {
					continue
				}
				if board[i][p.y] == '#' {
					return *p, false
				}
				return pos{
					x:         i,
					y:         p.y,
					direction: p.direction,
				}, true
			}
		}
	}
	if directions[p.direction].x == -1 {
		if nx < 0 || board[nx][ny] == ' ' {
			for i := len(board) - 1; i > 0; i-- {
				if board[i][p.y] == ' ' {
					continue
				}
				if board[i][p.y] == '#' {
					return *p, false
				}
				return pos{
					x:         i,
					y:         p.y,
					direction: p.direction,
				}, true
			}
		}
	}
	if directions[p.direction].y == 1 {
		if ny >= len(board[nx]) || board[nx][ny] == ' ' {
			for i := 0; i < len(board[nx]); i++ {
				if board[p.x][i] == ' ' {
					continue
				}
				if board[p.x][i] == '#' {
					return *p, false
				}
				return pos{
					x:         p.x,
					y:         i,
					direction: p.direction,
				}, true
			}
		}
	}
	if directions[p.direction].y == -1 {
		if ny < 0 || board[nx][ny] == ' ' {
			for i := len(board[nx]) - 1; i > 0; i-- {
				if board[p.x][i] == ' ' {
					continue
				}
				if board[p.x][i] == '#' {
					return *p, false
				}
				return pos{
					x:         p.x,
					y:         i,
					direction: p.direction,
				}, true
			}
		}
	}
	if board[nx][ny] == '#' {
		return *p, false
	}
	return p.simpleMove(), true
}

var (
	k = 50
)

func newPos3D(p *pos, board [][]rune) pos {
	nx := p.x + directions[p.direction].x
	ny := p.y + directions[p.direction].y
	nd := p.direction

	if p.direction == downI && (nx >= len(board) || board[nx][ny] == ' ') {
		switch p.y / k {
		case 0:
			nx = 0
			ny = 2*k + p.y
		case 1:
			nx = 3*k + (p.y % k)
			ny = k - 1
			nd = leftI
		case 2:
			nx = k + (p.y % k)
			ny = 2*k - 1
			nd = leftI
		}
	} else if p.direction == upI && (nx < 0 || board[nx][ny] == ' ') {
		switch p.y / k {
		case 0:
			nx = k + p.y
			ny = k
			nd = rightI
		case 1:
			nx = 3*k + (p.y % k)
			ny = 0
			nd = rightI
		case 2:
			nx = 4*k - 1
			ny = p.y % k
		}
	} else if p.direction == rightI && (ny >= len(board[nx]) || board[nx][ny] == ' ') {
		switch p.x / k {
		case 0:
			nx = 3*k - 1 - (p.x % k)
			ny = 2*k - 1
			nd = leftI
		case 1:
			nx = k - 1
			ny = 2*k + (p.x % k)
			nd = upI
		case 2:
			nx = k - 1 - (p.x % k)
			ny = 3*k - 1
			nd = leftI
		case 3:
			nx = 3*k - 1
			ny = k + (p.x % k)
			nd = upI
		}
	} else if p.direction == leftI && (ny < 0 || board[nx][ny] == ' ') {
		switch p.x / k {
		case 0:
			nx = 3*k - 1 - (p.x % k)
			ny = 0
			nd = rightI
		case 1:
			nx = 2 * k
			ny = p.x % k
			nd = downI
		case 2:
			nx = k - 1 - (p.x % k)
			ny = k
			nd = rightI
		case 3:
			nx = 0
			ny = k + (p.x % k)
			nd = downI
		}
	}
	return pos{nx, ny, nd}
}

func (p *pos) move3D(board [][]rune) (pos, bool) {
	n := newPos3D(p, board)

	if board[n.x][n.y] == '#' {
		return *p, false
	}

	return n, true
}

func (p *pos) part() int {
	a := (1 + p.x) / 50
	b := (1 + p.y) / 50
	return 5*a + b
}

func (p *pos) simpleMove() pos {
	return pos{
		x:         p.x + directions[p.direction].x,
		y:         p.y + directions[p.direction].y,
		direction: p.direction,
	}
}

type direction struct {
	x, y int
}

type move struct {
	forward  int
	rotation int
}

func parseMoves(s string) []*move {
	var moves []*move
	acc := ""
	for _, c := range s {
		if c == 'R' {
			forward, _ := strconv.Atoi(acc)
			moves = append(moves, &move{forward: forward})
			moves = append(moves, &move{rotation: +1})
			acc = ""
		} else if c == 'L' {
			forward, _ := strconv.Atoi(acc)
			moves = append(moves, &move{forward: forward})
			moves = append(moves, &move{rotation: -1})
			acc = ""
		} else {
			acc += string(c)
		}
	}
	if len(acc) > 0 {
		forward, _ := strconv.Atoi(acc)
		moves = append(moves, &move{forward: forward})
	}
	return moves
}

func partTwo(s string) int {
	m, moves, ok := strings.Cut(s, "\n\n")
	if !ok {
		panic("not ok")
	}
	var board [][]rune
	maxWidth := 0
	for _, l := range strings.Split(m, "\n") {
		var boardLine []rune
		for _, c := range l {
			boardLine = append(boardLine, c)
		}
		if len(boardLine) > maxWidth {
			maxWidth = len(boardLine)
		}
		board = append(board, boardLine)
	}

	for i := range board {
		for len(board[i]) < maxWidth {
			board[i] = append(board[i], ' ')
		}
	}

	verify(board)

	p := pos{0, findStartY(board), 0}
	for _, m := range parseMoves(moves) {
		if m.rotation != 0 {
			p.direction += m.rotation
			if p.direction < 0 {
				p.direction = len(directions) - 1
			}
			p.direction = p.direction % len(directions)
		} else {
			for i := 0; i < m.forward; i++ {
				moved := false
				p, moved = p.move3D(board)
				if !moved {
					break
				}
			}
		}
	}

	return 1000*(p.x+1) + 4*(p.y+1) + p.direction
}

func verify(board [][]rune) {
	var exp, got pos
	// 1 -> 6
	exp, got = pos{3 * k, 0, rightI}, newPos3D(&pos{0, k, upI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{0, k, downI}, newPos3D(&pos{3 * k, 0, leftI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 1 -> 4
	exp, got = pos{3*k - 1, 0, rightI}, newPos3D(&pos{0, k, leftI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{0, k, rightI}, newPos3D(&pos{3*k - 1, 0, leftI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 2 -> 6
	exp, got = pos{4*k - 1, k - 1, upI}, newPos3D(&pos{0, 3*k - 1, upI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{0, 3*k - 1, downI}, newPos3D(&pos{4*k - 1, k - 1, downI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 2 -> 5
	exp, got = pos{3*k - 1, 2*k - 1, leftI}, newPos3D(&pos{0, 3*k - 1, rightI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{0, 3*k - 1, leftI}, newPos3D(&pos{3*k - 1, 2*k - 1, rightI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 2 -> 3
	exp, got = pos{2*k - 1, 2*k - 1, leftI}, newPos3D(&pos{k - 1, 3*k - 1, downI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{k - 1, 3*k - 1, upI}, newPos3D(&pos{2*k - 1, 2*k - 1, rightI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 3 -> 4
	exp, got = pos{2 * k, 0, downI}, newPos3D(&pos{k, k, leftI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{k, k, rightI}, newPos3D(&pos{2 * k, 0, upI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}

	// 5 -> 6
	exp, got = pos{4*k - 1, k - 1, leftI}, newPos3D(&pos{3*k - 1, 2*k - 1, downI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
	exp, got = pos{3*k - 1, 2*k - 1, upI}, newPos3D(&pos{4*k - 1, k - 1, rightI}, board)
	if got != exp {
		log.Fatalf("expected: %v, got: %v", exp, got)
	}
}
