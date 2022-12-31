package main

import (
	"log"
	"os"
	"strings"
	"time"
)

const (
	test = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
)

func main() {
	data, err := os.ReadFile("./advent2022/25/input.txt")
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

type snafu struct {
	v []int
}

func (s *snafu) value() int {
	v := 0
	p := 1
	for i := len(s.v) - 1; i >= 0; i-- {
		v += s.v[i] * p
		p = p * 5
	}
	return v
}

func (s *snafu) String() string {
	var result string
	for i := len(s.v) - 1; i >= 0; i-- {
		switch s.v[i] {
		case 2:
			result += "2"
		case 1:
			result += "1"
		case 0:
			result += "0"
		case -1:
			result += "-"
		case -2:
			result += "="
		}
	}
	return result
}

func newSnafu(i int) *snafu {
	var v []int
	for i > 0 {
		i += 2
		k := i % 5
		i = i / 5
		v = append(v, k-2)
	}
	return &snafu{v}
}

func partOne(s string) string {
	var snafus []*snafu
	for _, l := range strings.Split(s, "\n") {
		var v []int
		for _, c := range l {
			switch c {
			case '=':
				v = append(v, -2)
			case '-':
				v = append(v, -1)
			case '0':
				v = append(v, 0)
			case '1':
				v = append(v, 1)
			case '2':
				v = append(v, 2)
			}
		}
		snafus = append(snafus, &snafu{v})
	}

	sum := 0
	for _, sn := range snafus {
		sum += sn.value()
	}

	log.Printf("sum: %v", sum)

	return newSnafu(sum).String()
}

func partTwo(s string) int {
	return 0
}
