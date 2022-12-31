package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	test = `1
2
-3
3
-2
0
4`
)

func main() {
	data, err := os.ReadFile("./advent2022/20/input.txt")
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

type val struct {
	v           int
	left, right *val
}

func partOne(s string) int {
	var values []*val
	var inOrder []*val
	var zeroValue *val
	for _, l := range strings.Split(s, "\n") {
		i, _ := strconv.Atoi(l)
		v := &val{
			v: i,
		}
		if len(values) > 0 {
			v.left = values[len(values)-1]
			values[len(values)-1].right = v
		}
		values = append(values, v)
		inOrder = append(inOrder, v)
		if v.v == 0 {
			zeroValue = v
		}
	}

	values[len(values)-1].right = values[0]
	values[0].left = values[len(values)-1]

	for _, v := range inOrder {
		if v.v > 0 {
			for i := 0; i < v.v; i++ {
				right := v.right
				left := v.left
				left.right = right
				right.left = left

				v.right = right.right
				right.right.left = v
				v.left = right
				right.right = v
			}
		} else if v.v < 0 {
			for i := 0; i < -v.v; i++ {
				right := v.right
				left := v.left
				left.right = right
				right.left = left

				v.left = left.left
				left.left.right = v
				v.right = left
				left.left = v
			}
		}
	}

	v := zeroValue
	sum := 0
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	sum += v.v
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	sum += v.v
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	sum += v.v
	return sum
}

func partTwo(s string) int {
	var values []*val
	var zeroValue *val
	for _, l := range strings.Split(s, "\n") {
		i, _ := strconv.Atoi(l)
		v := &val{
			v: i * 811589153,
		}
		if len(values) > 0 {
			v.left = values[len(values)-1]
			values[len(values)-1].right = v
		}
		values = append(values, v)
		if v.v == 0 {
			zeroValue = v
		}
	}

	values[len(values)-1].right = values[0]
	values[0].left = values[len(values)-1]

	for i := 0; i < 10; i++ {
		for _, v := range values {
			if v.v > 0 {
				k := v.v % (len(values) - 1)
				for i := 0; i < k; i++ {
					right := v.right
					left := v.left
					left.right = right
					right.left = left

					v.right = right.right
					right.right.left = v
					v.left = right
					right.right = v
				}
			} else if v.v < 0 {
				k := (-v.v) % (len(values) - 1)
				for i := 0; i < k; i++ {
					right := v.right
					left := v.left
					left.right = right
					right.left = left

					v.left = left.left
					left.left.right = v
					v.right = left
					left.left = v
				}
			}
		}
	}

	v := zeroValue
	sum := 0
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	log.Printf("1000th: %v", v.v)
	sum += v.v
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	log.Printf("2000th: %v", v.v)
	sum += v.v
	for i := 0; i < 1000; i++ {
		v = v.right
	}
	log.Printf("3000th: %v", v.v)
	sum += v.v
	return sum
}
