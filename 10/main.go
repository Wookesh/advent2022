package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const test = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

func main() {
	data, err := os.ReadFile("./advent2022/10/input.txt")
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

type CPU struct {
	X int
}

type Instruction interface {
	Cycles() int
	Exec(*CPU) bool
}

type NoOp struct {
}

func (n *NoOp) Cycles() int {
	return 1
}

func (n *NoOp) Exec(c *CPU) bool { return true }

func (n *NoOp) String() string {
	return "noop"
}

type AddX struct {
	x int
}

func NewAddX(x int) *AddX {
	return &AddX{
		x: x,
	}
}

func (a *AddX) Cycles() int {
	return 2
}

func (a *AddX) Exec(c *CPU) bool {
	c.X += a.x
	return true
}

func (a *AddX) String() string {
	return fmt.Sprintf("addx %v", a.x)
}

func partOne(s string) int {
	var instructions []Instruction
	for _, l := range strings.Split(s, "\n") {
		if l == "noop" {
			instructions = append(instructions, &NoOp{})
		} else {
			var x int
			_, err := fmt.Sscanf(l, "addx %v", &x)
			if err != nil {
				log.Fatal(err)
			}
			instructions = append(instructions, NewAddX(x))
		}
	}

	cpu := &CPU{X: 1}

	currentInstructionIndex := 0
	currentInstructionCycles := 0
	result := 0
	for i := 1; i <= 220; i++ {
		if (i+20)%40 == 0 {
			signalStrength := i * cpu.X
			result += signalStrength
		}
		currentInstructionCycles += 1
		if currentInstructionCycles == instructions[currentInstructionIndex].Cycles() {
			instructions[currentInstructionIndex].Exec(cpu)
			currentInstructionCycles = 0
			currentInstructionIndex = (currentInstructionIndex + 1) % len(instructions)
		}
	}

	return result
}

func partTwo(s string) int {
	var instructions []Instruction
	for _, l := range strings.Split(s, "\n") {
		if l == "noop" {
			instructions = append(instructions, &NoOp{})
		} else {
			var x int
			_, err := fmt.Sscanf(l, "addx %v", &x)
			if err != nil {
				log.Fatal(err)
			}
			instructions = append(instructions, NewAddX(x))
		}
	}

	cpu := &CPU{X: 1}

	currentInstructionIndex := 0
	currentInstructionCycles := 0
	result := 0
	for i := 1; i <= 240; i++ {
		//log.Printf("cycle %v: x: %v, op: %v", i, cpu.X, instructions[currentInstructionIndex])
		if (i+20)%40 == 0 {
			signalStrength := i * cpu.X
			//log.Printf("cycle %v: x: %v, signalStreng: %v", i, cpu.X, signalStrength)
			result += signalStrength
		}
		if cpu.X <= i%40 && i%40 <= cpu.X+2 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if i%40 == 0 {
			fmt.Print("\n")
		}
		currentInstructionCycles += 1
		if currentInstructionCycles == instructions[currentInstructionIndex].Cycles() {
			instructions[currentInstructionIndex].Exec(cpu)
			currentInstructionCycles = 0
			currentInstructionIndex = (currentInstructionIndex + 1) % len(instructions)
		}
	}

	return result

}
