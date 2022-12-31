package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wookesh/advent2022/07/fs"
)

const test = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func main() {
	data, err := os.ReadFile("./advent2022/07/input.txt")
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

const (
	ls   = "ls"
	cd   = "cd"
	none = ""
)

func partOne(s string) int {
	currentCMD := none
	fileSystem := fs.NewFS()
	for _, l := range strings.Split(s, "\n") {
		if strings.HasPrefix(l, "$") {
			cmd := strings.Split(strings.TrimPrefix(l, "$ "), " ")
			switch cmd[0] {
			case "ls":
				currentCMD = ls
			case "cd":
				fileSystem.CD(cmd[1])
			}
		} else {
			switch currentCMD {
			case ls:
				parts := strings.Split(l, " ")
				if parts[0] == "dir" {
					fileSystem.CurrentDir().Add(fs.NewDir(parts[1], fileSystem.CurrentDir()))
				} else {
					sizeS, name := parts[0], parts[1]
					size, _ := strconv.Atoi(sizeS)
					fileSystem.CurrentDir().Add(fs.NewFile(name, size))
				}
			}
		}
	}

	totalSize := 0
	fileSystem.Walk(func(e fs.Entity) {
		if !e.IsDir() {
			return
		}
		if e.Size() < 100000 {
			totalSize += e.Size()
		}
	})

	return totalSize
}

func partTwo(s string) int {
	currentCMD := none
	fileSystem := fs.NewFS()
	for _, l := range strings.Split(s, "\n") {
		if strings.HasPrefix(l, "$") {
			cmd := strings.Split(strings.TrimPrefix(l, "$ "), " ")
			switch cmd[0] {
			case "ls":
				currentCMD = ls
			case "cd":
				fileSystem.CD(cmd[1])
			}
		} else {
			switch currentCMD {
			case ls:
				parts := strings.Split(l, " ")
				if parts[0] == "dir" {
					fileSystem.CurrentDir().Add(fs.NewDir(parts[1], fileSystem.CurrentDir()))
				} else {
					sizeS, name := parts[0], parts[1]
					size, _ := strconv.Atoi(sizeS)
					fileSystem.CurrentDir().Add(fs.NewFile(name, size))
				}
			}
		}
	}

	max := 70000000
	required := 30000000
	current := fileSystem.Root().Size()

	min := current

	fileSystem.Walk(func(e fs.Entity) {
		if !e.IsDir() {
			return
		}
		newFree := max - current + e.Size()
		if newFree < required {
			return
		}
		if min > e.Size() {
			min = e.Size()
		}
	})

	return min
}
