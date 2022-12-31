package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const test = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func main() {
	data, err := os.ReadFile("./advent2022/11/input.txt")
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

type Monkey struct {
	Items        []int
	BigItems     []*big.Int
	Operation    func(int) int
	BigOperation func(*big.Int) *big.Int
	Test         int
	IfTrue       int
	IfFalse      int
}

func partOne(s string) int {
	var monkeys []*Monkey
	for _, desc := range strings.Split(s, "\n\n") {
		var m Monkey
		parts := strings.Split(desc, "\n")
		startingItems := strings.Split(strings.TrimPrefix(parts[1], "  Starting items: "), ", ")
		for _, si := range startingItems {
			i, _ := strconv.Atoi(si)
			m.Items = append(m.Items, i)
		}
		var multiplier string
		var op string
		fmt.Sscanf(parts[2], "  Operation: new = old %v %v", &op, &multiplier)
		m.Operation = parseOp(multiplier, op)
		fmt.Sscanf(parts[3], "  Test: divisible by %v", &m.Test)
		fmt.Sscanf(parts[4], "    If true: throw to monkey %v", &m.IfTrue)
		fmt.Sscanf(parts[5], "    If false: throw to monkey %v", &m.IfFalse)
		//log.Info(m)
		monkeys = append(monkeys, &m)
	}

	counters := make([]int, len(monkeys))
	for i := 1; i <= 20; i++ {
		for mi, m := range monkeys {
			items := m.Items
			m.Items = nil
			for _, item := range items {
				counters[mi]++
				//log.Printf("Monkey inspects an item with a worry level of %v.", item)
				worry := m.Operation(item)
				//log.Printf("  Worry level is multiplied to %v.", worry)
				worry = worry / 3
				//log.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %v.", worry)
				if worry%m.Test == 0 {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfTrue)
					monkeys[m.IfTrue].Items = append(monkeys[m.IfTrue].Items, worry)
				} else {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfFalse)
					monkeys[m.IfFalse].Items = append(monkeys[m.IfFalse].Items, worry)
				}
			}
		}

		for mi, m := range monkeys {
			log.Printf("Round: %v, Monkey: %v, items: %v", i, mi, m.Items)
		}
	}

	sort.Ints(counters)

	return counters[len(counters)-1] * counters[len(counters)-2]
}
func partTwo(s string) int {
	var monkeys []*Monkey
	for _, desc := range strings.Split(s, "\n\n") {
		var m Monkey
		parts := strings.Split(desc, "\n")
		startingItems := strings.Split(strings.TrimPrefix(parts[1], "  Starting items: "), ", ")
		for _, si := range startingItems {
			i, _ := strconv.Atoi(si)
			m.Items = append(m.Items, i)
		}
		var multiplier string
		var op string
		fmt.Sscanf(parts[2], "  Operation: new = old %v %v", &op, &multiplier)
		m.Operation = parseOp(multiplier, op)
		fmt.Sscanf(parts[3], "  Test: divisible by %v", &m.Test)
		fmt.Sscanf(parts[4], "    If true: throw to monkey %v", &m.IfTrue)
		fmt.Sscanf(parts[5], "    If false: throw to monkey %v", &m.IfFalse)
		//log.Info(m)
		monkeys = append(monkeys, &m)
	}

	lcm := 1
	for _, m := range monkeys {
		//lcm = lcm * m.Test
		lcm = lowestCommonMultiple(lcm, m.Test)
	}

	counters := make([]int, len(monkeys))
	for i := 1; i <= 10000; i++ {
		for mi, m := range monkeys {
			items := m.Items
			m.Items = nil
			for _, item := range items {
				counters[mi]++
				//log.Printf("Monkey inspects an item with a worry level of %v.", item)
				worry := m.Operation(item)
				//log.Printf("  Worry level is multiplied to %v.", worry)
				//worry = worry / 3
				//if worry%lcm == 0 {
				worry = worry % lcm
				//}
				//log.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %v.", worry)
				if worry%m.Test == 0 {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfTrue)
					monkeys[m.IfTrue].Items = append(monkeys[m.IfTrue].Items, worry)
				} else {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfFalse)
					monkeys[m.IfFalse].Items = append(monkeys[m.IfFalse].Items, worry)
				}
			}
		}

		if i%1000 == 0 {
			for mi, m := range monkeys {
				log.Printf("Round: %v, Monkey: %v, items: %v", i, mi, m.Items)
			}
		}
	}

	sort.Ints(counters)

	return counters[len(counters)-1] * counters[len(counters)-2]
}

func parseOp(x string, op string) func(int) int {
	var f func(int, int) int
	switch op {
	case "+":
		f = func(a, b int) int { return a + b }
	case "*":
		f = func(a, b int) int { return a * b }
	}
	if x == "old" {
		return func(old int) int {
			return f(old, old)
		}
	}
	xi, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	return func(old int) int {
		return f(old, xi)
	}
}

func parseOpBig(x string, op string) func(*big.Int) *big.Int {
	var f func(*big.Int, *big.Int) *big.Int
	switch op {
	case "+":
		f = func(a, b *big.Int) *big.Int { z := big.NewInt(0); return z.Add(a, b) }
	case "*":
		f = func(a, b *big.Int) *big.Int { z := big.NewInt(0); return z.Mul(a, b) }
	}
	if x == "old" {
		return func(old *big.Int) *big.Int {
			return f(old, old)
		}
	}
	xi, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	return func(old *big.Int) *big.Int {
		return f(old, big.NewInt(int64(xi)))
	}
}

func partTwoBig(s string) int {
	var monkeys []*Monkey
	for _, desc := range strings.Split(s, "\n\n") {
		var m Monkey
		parts := strings.Split(desc, "\n")
		startingItems := strings.Split(strings.TrimPrefix(parts[1], "  Starting items: "), ", ")
		for _, si := range startingItems {
			i, _ := strconv.Atoi(si)
			m.Items = append(m.Items, i)
		}
		var multiplier string
		var op string
		fmt.Sscanf(parts[2], "  Operation: new = old %v %v", &op, &multiplier)
		//m.Operation = parseOp(multiplier, op)
		m.BigOperation = parseOpBig(multiplier, op)
		fmt.Sscanf(parts[3], "  Test: divisible by %v", &m.Test)
		fmt.Sscanf(parts[4], "    If true: throw to monkey %v", &m.IfTrue)
		fmt.Sscanf(parts[5], "    If false: throw to monkey %v", &m.IfFalse)
		//log.Info(m)
		monkeys = append(monkeys, &m)
	}

	for _, m := range monkeys {
		for _, i := range m.Items {
			m.BigItems = append(m.BigItems, big.NewInt(int64(i)))
		}
	}

	counters := make([]int, len(monkeys))
	for i := 1; i <= 10000; i++ {
		for mi, m := range monkeys {
			items := m.BigItems
			m.BigItems = nil
			for _, item := range items {
				counters[mi]++
				//log.Printf("Monkey inspects an item with a worry level of %v.", item)
				worry := m.BigOperation(item)
				//log.Printf("  Worry level is multiplied to %v.", worry)
				//worry = worry / 3
				//log.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %v.", worry)
				z := big.NewInt(0)
				z.Mod(worry, big.NewInt(int64(m.Test)))
				if z.Cmp(big.NewInt(0)) == 0 {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfTrue)
					monkeys[m.IfTrue].BigItems = append(monkeys[m.IfTrue].BigItems, worry)
				} else {
					//log.Printf("  Item with worry level %v is thrown to monkey %v.", worry, m.IfFalse)
					monkeys[m.IfFalse].BigItems = append(monkeys[m.IfFalse].BigItems, worry)
				}
			}
		}

		if i%100 == 0 {
			for mi, m := range monkeys {
				log.Printf("Round: %v, Monkey: %v, items: %v", i, mi, m.Items)
			}
			log.Print(counters)
		}
	}

	sort.Ints(counters)

	return counters[len(counters)-1] * counters[len(counters)-2]
}

func greatestCommonDivisor(u, v int) int {
	if v > u {
		u, v = v, u
	}
	for v != 0 {
		u, v = v, u%v
	}
	return u
}

func lowestCommonMultiple(u, v int) int {
	gcd := greatestCommonDivisor(u, v)
	return (u / gcd) * v
}
