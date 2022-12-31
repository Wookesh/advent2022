package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func main() {
	data, err := os.ReadFile("./advent2022/05/input.txt")
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

type op struct {
	Count int
	Src   int
	Dst   int
}

func partTwo(input string) string {
	var stacks [][]byte
	var operations []op

	for _, l := range strings.Split(input, "\n") {
		if stacks == nil {
			stacks = make([][]byte, len(l)/4+1)
		}
		if strings.ContainsRune(l, '[') {
			for i := 0; ; i += 4 {
				si := i / 4
				if si*4 >= len(l) {
					break
				}
				if l[i+1] == ' ' {
					continue
				}
				stacks[si] = append(stacks[si], l[i+1])
				fmt.Printf("%v, %v\n", si, string(stacks[si]))
			}
		} else {
			var operation op
			_, err := fmt.Sscanf(l, "move %v from %v to %v", &(operation.Count), &(operation.Src), &(operation.Dst))
			if err != nil {
				continue
			}
			operations = append(operations, operation)
		}
	}

	for i := range stacks {
		stacks[i] = reverse(stacks[i])
	}

	for _, o := range operations {
		src := o.Src - 1
		dst := o.Dst - 1

		v := stacks[src][len(stacks[src])-o.Count:]
		stacks[src] = stacks[src][:len(stacks[src])-o.Count]
		stacks[dst] = append(stacks[dst], v...)
	}

	var result string
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}

func partOne(input string) string {
	var stacks [][]byte
	var operations []op

	for _, l := range strings.Split(input, "\n") {
		if stacks == nil {
			stacks = make([][]byte, len(l)/4+1)
		}
		if strings.ContainsRune(l, '[') {
			for i := 0; ; i += 4 {
				si := i / 4
				if si*4 >= len(l) {
					break
				}
				if l[i+1] == ' ' {
					continue
				}
				stacks[si] = append(stacks[si], l[i+1])
				fmt.Printf("%v, %v\n", si, string(stacks[si]))
			}
		} else {
			var operation op
			_, err := fmt.Sscanf(l, "move %v from %v to %v", &(operation.Count), &(operation.Src), &(operation.Dst))
			if err != nil {
				continue
			}
			operations = append(operations, operation)
		}
	}

	for i := range stacks {
		stacks[i] = reverse(stacks[i])

		fmt.Printf("%v, %v\n", i, stacks[i])
	}

	for _, o := range operations {
		fmt.Println(o)
		src := o.Src - 1
		dst := o.Dst - 1
		for i := 0; i < o.Count; i++ {
			v := stacks[src][len(stacks[src])-1]
			stacks[dst] = append(stacks[dst], v)
			stacks[src] = stacks[src][:len(stacks[src])-1]
		}
	}

	var result string
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}

func reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i := range items {
		result[i] = items[len(items)-1-i]
	}
	return result
}
