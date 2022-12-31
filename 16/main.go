package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func main() {
	data, err := os.ReadFile("./advent2022/16/input.txt")
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

type Valve struct {
	ID     string
	Rate   int
	Valves []string
}

func partOne(s string) int {
	valves := make(map[string]*Valve)
	for _, l := range strings.Split(s, "\n") {
		first, second, _ := strings.Cut(l, ";")
		var rate int
		var valve string
		_, err := fmt.Sscanf(first, "Valve %v has flow rate=%v", &valve, &rate)
		if err != nil {
			log.Fatal(err)
		}

		second = strings.TrimPrefix(second, " tunnels lead to valves ")
		second = strings.TrimPrefix(second, " tunnel leads to valve ")
		vs := strings.Split(second, ", ")
		valves[valve] = &Valve{
			ID:     valve,
			Rate:   rate,
			Valves: vs,
		}
	}

	distances := make(map[string]map[string]int)

	for _, v := range valves {
		distances[v.ID] = make(map[string]int)
		distances[v.ID][v.ID] = 0
		for _, next := range v.Valves {
			distances[v.ID][next] = 1
		}
	}

	for _, v := range valves {
		for _, u := range valves {
			if d, ok := distances[v.ID][u.ID]; v.ID == u.ID || ok {
				distances[u.ID][v.ID] = d
				continue
			} else {
				distances[u.ID][v.ID] = d
			}

			type item struct {
				vid  string
				dist int
			}

			queue := []item{{vid: v.ID, dist: 0}}
			seen := make(map[string]bool)

		MAIN:
			for len(queue) > 0 {
				i := queue[0]
				queue = queue[1:]
				seen[i.vid] = true

				q := valves[i.vid]

				for _, n := range q.Valves {
					if n == u.ID {
						distances[v.ID][u.ID] = i.dist + 1
						distances[u.ID][v.ID] = i.dist + 1
						break MAIN
					}
					if _, ok := seen[n]; ok {
						continue
					}
					seen[n] = true

					queue = append(queue, item{vid: n, dist: i.dist + 1})
				}
			}
		}
	}

	//for k, d := range distances {
	//	for l, di := range d {
	//		log.Printf("%v -> %v: %v", k, l, di)
	//	}
	//}

	return findBestOpening(valves, distances)
}

func findBestOpening(valves map[string]*Valve, distances map[string]map[string]int) int {
	start := valves["AA"]
	minutes := 30

	var closed []string
	for _, v := range valves {
		if v.Rate > 0 {
			closed = append(closed, v.ID)
		}
	}

	return findBestOpeningI(start, minutes, closed, valves, distances)
}

func findBestOpeningI(v *Valve, minutes int, closed []string, valves map[string]*Valve, distances map[string]map[string]int) int {
	max := 0
	for _, c := range closed {
		mLeft := minutes - distances[v.ID][c] - 1
		if mLeft <= 0 {
			continue
		}
		closed2 := copyWithout(closed, c)
		pressure := findBestOpeningI(valves[c], mLeft, closed2, valves, distances)
		mp := pressure + mLeft*valves[c].Rate
		if mp > max {
			max = mp
		}
	}
	return max
}

func copyWithout(l []string, s string) []string {
	var result []string
	for _, e := range l {
		if e == s {
			continue
		}
		result = append(result, e)
	}
	return result
}

func partTwo(s string) int {
	valves := make(map[string]*Valve)
	for _, l := range strings.Split(s, "\n") {
		first, second, _ := strings.Cut(l, ";")
		var rate int
		var valve string
		_, err := fmt.Sscanf(first, "Valve %v has flow rate=%v", &valve, &rate)
		if err != nil {
			log.Fatal(err)
		}

		second = strings.TrimPrefix(second, " tunnels lead to valves ")
		second = strings.TrimPrefix(second, " tunnel leads to valve ")
		vs := strings.Split(second, ", ")
		valves[valve] = &Valve{
			ID:     valve,
			Rate:   rate,
			Valves: vs,
		}
	}

	distances := make(map[string]map[string]int)

	for _, v := range valves {
		distances[v.ID] = make(map[string]int)
		distances[v.ID][v.ID] = 0
		for _, next := range v.Valves {
			distances[v.ID][next] = 1
		}
	}

	for _, v := range valves {
		for _, u := range valves {
			if d, ok := distances[v.ID][u.ID]; v.ID == u.ID || ok {
				distances[u.ID][v.ID] = d
				continue
			} else {
				distances[u.ID][v.ID] = d
			}

			type item struct {
				vid  string
				dist int
			}

			queue := []item{{vid: v.ID, dist: 0}}
			seen := make(map[string]bool)

		MAIN:
			for len(queue) > 0 {
				i := queue[0]
				queue = queue[1:]
				seen[i.vid] = true

				q := valves[i.vid]

				for _, n := range q.Valves {
					if n == u.ID {
						distances[v.ID][u.ID] = i.dist + 1
						distances[u.ID][v.ID] = i.dist + 1
						break MAIN
					}
					if _, ok := seen[n]; ok {
						continue
					}
					seen[n] = true

					queue = append(queue, item{vid: n, dist: i.dist + 1})
				}
			}
		}
	}

	//for k, d := range distances {
	//	for l, di := range d {
	//		log.Printf("%v -> %v: %v", k, l, di)
	//	}
	//}

	return findBestOpeningWithElephant(valves, distances)
}

func findBestOpeningWithElephant(valves map[string]*Valve, distances map[string]map[string]int) int {
	start := valves["AA"]
	minutes := 26

	var closed []string
	for _, v := range valves {
		if v.Rate > 0 {
			closed = append(closed, v.ID)
		}
	}

	max := 0
	for _, split := range getCombinations(closed) {
		a, b := split[0], split[1]
		ma := findBestOpeningI(start, minutes, a, valves, distances)
		mb := findBestOpeningI(start, minutes, b, valves, distances)
		if ma+mb > max {
			max = ma + mb
		}
	}

	return max
}

func getCombinations(list []string) [][][]string {
	var result [][][]string

	for i := 0; i < (1 << uint(len(list))); i++ {
		var part1, part2 []string
		for j, v := range list {
			if (i>>uint(j))&1 == 1 {
				part1 = append(part1, v)
			} else {
				part2 = append(part2, v)
			}
		}
		result = append(result, [][]string{part1, part2})
	}
	return result
}
