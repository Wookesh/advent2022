package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	test = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
)

func main() {
	data, err := os.ReadFile("./advent2022/21/input.txt")
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

type Value interface {
	Value() int
	HasHuman() bool
	FindHumanValue(k int) int
}

type SimpleValue struct {
	name  string
	v     int
	panic bool
}

func (v *SimpleValue) Value() int {
	if v.panic && v.name == "humn" {
		log.Fatalf("humn was called for value")
	}
	return v.v
}

func (v *SimpleValue) HasHuman() bool {
	return v.name == "humn"
}

func (v *SimpleValue) FindHumanValue(k int) int {
	if v.name != "humn" {
		log.Fatalf("impossbru: %v", v.name)
	}
	return k
}

type OperationValue struct {
	name string
	a, b string
	op   string
	m    map[string]Value
}

func (v *OperationValue) Value() int {
	ma, ok := v.m[v.a]
	if !ok {
		log.Fatalf("missing monkey: %v", v.a)
	}
	av := ma.Value()
	mb, ok := v.m[v.b]
	if !ok {
		log.Fatalf("missing monkey: %v", v.b)
	}

	bv := mb.Value()
	switch v.op {
	case "+":
		return av + bv
	case "-":
		return av - bv
	case "*":
		return av * bv
	case "/":
		return av / bv
	case "=":
		if av == bv {
			return 0
		} else if av < bv {
			return -1
		} else {
			return 1
		}
	default:
		log.Fatalf("unknown op: %v", v.op)
		return 0
	}
}

func (v *OperationValue) HasHuman() bool {
	ma, ok := v.m[v.a]
	if !ok {
		log.Fatalf("missing monkey: %v", v.a)
	}
	mb, ok := v.m[v.b]
	if !ok {
		log.Fatalf("missing monkey: %v", v.b)
	}
	return ma.HasHuman() || mb.HasHuman()
}

func (v *OperationValue) FindHumanValue(k int) int {
	ma, ok := v.m[v.a]
	if !ok {
		log.Fatalf("missing monkey: %v", v.a)
	}
	mb, ok := v.m[v.b]
	if !ok {
		log.Fatalf("missing monkey: %v", v.b)
	}

	switch v.op {
	case "=":
		if ma.HasHuman() {
			return ma.FindHumanValue(mb.Value())
		}
		if mb.HasHuman() {
			return mb.FindHumanValue(ma.Value())
		}
		return 0
	// a + b = k
	case "+":
		if ma.HasHuman() {
			return ma.FindHumanValue(k - mb.Value())
		}
		if mb.HasHuman() {
			return mb.FindHumanValue(k - ma.Value())
		}
	// a - b = k
	case "-":
		if ma.HasHuman() {
			return ma.FindHumanValue(k + mb.Value())
		}
		if mb.HasHuman() {
			return mb.FindHumanValue(ma.Value() - k)
		}
	// a * b = k
	case "*":
		if ma.HasHuman() {
			return ma.FindHumanValue(k / mb.Value())
		}
		if mb.HasHuman() {
			return mb.FindHumanValue(k / ma.Value())
		}
	// a / b = k
	case "/":
		if ma.HasHuman() {
			return ma.FindHumanValue(k * mb.Value())
		}
		if mb.HasHuman() {
			return mb.FindHumanValue(ma.Value() / k)
		}
	default:
		log.Fatalf("unknown op: %v", v.op)
		return 0
	}
	return 0
}

func partOne(s string) int {

	ops := make(map[string]Value)

	for _, l := range strings.Split(s, "\n") {
		result, rest, _ := strings.Cut(l, ":")
		rest = strings.TrimSpace(rest)
		val, err := strconv.Atoi(rest)
		if err != nil {
			parts := strings.Split(rest, " ")
			ops[result] = &OperationValue{
				name: result,
				a:    parts[0],
				b:    parts[2],
				op:   parts[1],
				m:    ops,
			}
		} else {
			ops[result] = &SimpleValue{name: result, v: val}
		}
	}
	return ops["root"].Value()
}

func partTwo(s string) int {
	ops := make(map[string]Value)

	for _, l := range strings.Split(s, "\n") {
		result, rest, _ := strings.Cut(l, ":")
		rest = strings.TrimSpace(rest)
		val, err := strconv.Atoi(rest)
		if err != nil {
			parts := strings.Split(rest, " ")
			ops[result] = &OperationValue{
				name: result,
				a:    parts[0],
				b:    parts[2],
				op:   parts[1],
				m:    ops,
			}
		} else {
			ops[result] = &SimpleValue{name: result, v: val, panic: true}
		}
	}

	ops["root"].(*OperationValue).op = "="

	return ops["root"].FindHumanValue(0)
}
