package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const test = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

type Item interface {
	Compare(other Item) (bool, bool)
	String() string
}

type List struct {
	Value   []Item
	divider bool
}

func (l *List) Compare(other Item) (correct bool, cont bool) {
	switch t := other.(type) {
	case *List:
		return compareLists(l, t)
	case *Number:
		return compareLists(l, &List{Value: []Item{t}})
	default:
		return false, false
	}
}

func compareLists(left, right *List) (bool, bool) {
	if len(left.Value) == 0 {
		return len(right.Value) >= 0, len(right.Value) == 0
	}
	for i := 0; ; i++ {
		if i >= len(left.Value) {
			return len(right.Value) >= len(left.Value), len(right.Value) <= len(left.Value)
		}
		if i >= len(right.Value) {
			return false, false
		}
		correct, cont := left.Value[i].Compare(right.Value[i])
		if !correct || !cont {
			return correct, cont
		}
	}
}

func (l *List) String() string {
	var parts []string
	for _, v := range l.Value {
		parts = append(parts, v.String())
	}

	return fmt.Sprintf("[%v]", strings.Join(parts, ","))
}

type Number struct {
	Value int
}

func (l *Number) Compare(other Item) (correct bool, cont bool) {
	switch t := other.(type) {
	case *Number:
		if l.Value < t.Value {
			return true, false
		} else if l.Value > t.Value {
			return false, false
		} else {
			return true, true
		}
	case *List:
		return compareLists(&List{Value: []Item{l}}, t)
	default:
		return false, false
	}
}

func (l *Number) String() string {
	return strconv.Itoa(l.Value)
}

func main() {
	data, err := os.ReadFile("./advent2022/13/input.txt")
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

type pair struct {
	left, right *List
}

func partOne(s string) int {
	var pairs []*pair
	for _, part := range strings.Split(s, "\n\n") {
		packages := strings.Split(part, "\n")

		left, err := parse(packages[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := parse(packages[1])
		if err != nil {
			log.Fatal(err)
		}

		pairs = append(pairs, &pair{left: left, right: right})
	}

	sum := 0
	for i, p := range pairs {
		if correct, _ := p.left.Compare(p.right); !correct {
			continue
		}
		log.Printf("correct: %v", i+1)
		sum += i + 1
	}

	return sum
}

func parse(s string) (*List, error) {
	list, rest, err := parseList([]byte(s)[1:])
	if err != nil {
		return nil, fmt.Errorf("parse(%v) failed rest: %v, err: %v", s, rest, err)
	}
	return list, nil
}

func parseList(line []byte) (*List, []byte, error) {
	var l List
	var err error
	for len(line) > 0 {
		c := line[0]
		switch c {
		case '[':
			var list *List
			line = line[1:]
			list, line, err = parseList(line)
			if err != nil {
				return nil, nil, err
			}
			l.Value = append(l.Value, list)
		case ']':
			line = line[1:]
			return &l, line, nil
		case ',':
			line = line[1:]
			continue
		default:
			if unicode.IsNumber(rune(c)) {
				var number *Number
				number, line, err = parseNumber(line)
				if err != nil {
					return nil, line, fmt.Errorf("parseNumber() failed: %w", err)
				}
				l.Value = append(l.Value, number)
			} else {
				return nil, nil, fmt.Errorf("unexpected character: '%v'", string(c))
			}
		}
	}
	return &l, line, nil
}

func parseNumber(line []byte) (*Number, []byte, error) {
	var number []byte
	for len(line) > 0 {
		c := line[0]
		if unicode.IsNumber(rune(c)) {
			number = append(number, c)
			line = line[1:]
		} else if c == ',' || c == '[' || c == ']' {
			break
		} else {
			return nil, nil, fmt.Errorf("unexpected character: '%v'", string(c))
		}
	}
	v, err := strconv.Atoi(string(number))
	if err != nil {
		return nil, line, err
	}
	return &Number{Value: v}, line, nil
}

type Lists struct {
	l []*List
}

func (l *Lists) Len() int {
	return len(l.l)
}

func (l *Lists) Less(i, j int) bool {
	ok, _ := l.l[i].Compare(l.l[j])
	return ok
}

func (l *Lists) Swap(i, j int) {
	l.l[i], l.l[j] = l.l[j], l.l[i]
}

func partTwo(s string) int {
	var packages Lists

	div1, err := parse("[[2]]")
	if err != nil {
		log.Fatal(err)
	}
	div1.divider = true
	div2, err := parse("[[6]]")
	if err != nil {
		log.Fatal(err)
	}
	div2.divider = true
	packages.l = append(packages.l, div1)
	packages.l = append(packages.l, div2)

	for _, part := range strings.Split(s, "\n") {
		if part == "" {
			continue
		}

		parsed, err := parse(part)
		if err != nil {
			log.Fatal(err)
		}
		packages.l = append(packages.l, parsed)
	}

	sort.Sort(&packages)

	sum := 1
	for i, l := range packages.l {
		if l.divider {
			sum *= i + 1
		}
	}

	return sum
}
